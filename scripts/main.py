import numpy as np
from pymongo import MongoClient
import os
import matplotlib.pyplot as plt


def remove_adjacent_duplicates(x, y):
    # Ensure both arrays are of the same length
    if len(x) != len(y):
        raise ValueError("Arrays x and y must have the same length")

    # Create a new list to store the filtered elements
    filtered_x = []
    filtered_y = []

    # Iterate through the arrays and filter out adjacent duplicates
    prev_y = None
    for index, (xi, yi) in enumerate(zip(x, y)):
        if yi != prev_y or index == len(x) - 1:
            filtered_x.append(xi)
            filtered_y.append(yi)
        prev_y = yi

    return filtered_x, filtered_y


def get_trainer_token_data(workspace_name):
    client = MongoClient("mongodb://127.0.0.1:27017")
    db = client["scatter-protocol"]

    workspace = db["workspaces"].find_one({"name": workspace_name})
    workspace_id = workspace["_id"]
    trainers = db["protocolnodes"].find(
        {"workspaceId": workspace_id, "peerType": "trainer"}
    )
    trainer_data = {}
    for index, trainer in enumerate(trainers):
        events = db["logevents"].find(
            {
                "blockchainAddress": trainer["blockchainAddress"],
                "workspaceId": workspace_id,
                "logType": "Token Balance",
            }
        )

        xData, yData = [], []
        for event in events:
            xData.append(event["xDataPoint"])
            yData.append(event["yDataPoint"])
        trainer_data[trainer["blockchainAddress"]] = {
            "x": xData,
            "y": yData,
            "type": "Trainer",
        }

    return trainer_data


def get_validator_token_data(workspace_name):
    client = MongoClient("mongodb://127.0.0.1:27017")
    db = client["scatter-protocol"]

    workspace = db["workspaces"].find_one({"name": workspace_name})
    workspace_id = workspace["_id"]
    validators = db["protocolnodes"].find(
        {"workspaceId": workspace_id, "peerType": "validator"}
    )
    validator_data = {}
    for index, validator in enumerate(validators):
        events = db["logevents"].find(
            {
                "blockchainAddress": validator["blockchainAddress"],
                "workspaceId": workspace_id,
                "logType": "Token Balance",
            }
        )

        xData, yData = [], []
        for event in events:
            xData.append(event["xDataPoint"])
            yData.append(event["yDataPoint"])
        validator_data[validator["blockchainAddress"]] = {
            "x": xData,
            "y": yData,
            "type": "Validator",
        }

    return validator_data


def get_challenger_token_data(workspace_name):
    client = MongoClient("mongodb://127.0.0.1:27017")
    db = client["scatter-protocol"]

    workspace = db["workspaces"].find_one({"name": workspace_name})
    workspace_id = workspace["_id"]
    validators = db["protocolnodes"].find(
        {"workspaceId": workspace_id, "peerType": "challenger"}
    )
    validator_data = {}
    for index, validator in enumerate(validators):
        events = db["logevents"].find(
            {
                "blockchainAddress": validator["blockchainAddress"],
                "workspaceId": workspace_id,
                "logType": "Token Balance",
            }
        )

        xData, yData = [], []
        for event in events:
            xData.append(event["xDataPoint"])
            yData.append(event["yDataPoint"])
        validator_data[validator["blockchainAddress"]] = {
            "x": xData,
            "y": yData,
            "type": "Challenger",
        }

    return validator_data


def create_overlayed_graph(
    data,
    output_file,
    x_axis_label,
    y_axis_label,
    fig_size,
    title,
    label_prefix=None,
    base_line=None,
    base_line_label="",
    secondary_base_line=None,
    secondary_base_line_label="",
):
    fig, ax = plt.subplots(figsize=fig_size)
    trainer_count = 0
    validator_count = 0

    for index, (key, values) in enumerate(data.items()):
        if values["type"] == "Trainer":
            trainer_count += 1
        if values["type"] == "Validator":
            validator_count += 1

        x_values = values["x"]
        y_values = values["y"]

        x_values, y_values = remove_adjacent_duplicates(x_values, y_values)

        node_idx = index + 1
        if values["type"] == "Validator":
            node_idx = index + 1 - trainer_count
        elif values["type"] == "Challenger":
            node_idx = index + 1 - trainer_count - validator_count

        ax.plot(
            x_values,
            y_values,
            label=f"{label_prefix if label_prefix else values['type']} {node_idx}",
        )

    ax.set_xlabel(x_axis_label)
    ax.set_ylabel(y_axis_label)
    if base_line:
        ax.axhline(y=base_line, color="red", linestyle="--", label=base_line_label)

    if secondary_base_line:
        ax.axhline(
            y=secondary_base_line,
            color="blue",
            linestyle="--",
            label=secondary_base_line_label,
        )

    ax.legend(
        fontsize=8,
        loc="upper center",
        bbox_to_anchor=(0.5, -0.08),
        fancybox=True,
        shadow=True,
        ncol=5,
    )

    plt.title(title)
    plt.tight_layout()
    plt.savefig(output_file)


def create_separated_graph(
    data,
    output_file,
    num_rows,
    num_cols,
    x_axis_label,
    y_axis_label,
    fig_size,
    title,
    label_prefix=None,
    subtitle_suffix="",
    base_line=None,
    base_line_label="",
    secondary_base_line=None,
    secondary_base_line_label="",
):
    fig, axes = plt.subplots(nrows=num_rows, ncols=num_cols, figsize=fig_size)
    if len(axes.shape) == 1:
        axes = np.array([axes])

    trainer_count = 0
    validator_count = 0
    for idx, (key, values) in enumerate(data.items()):
        if values["type"] == "Trainer":
            trainer_count += 1

        if values["type"] == "Validator":
            validator_count += 1

        row_idx = idx // num_cols
        col_idx = idx % num_cols

        x_values = values["x"]
        y_values = values["y"]

        x_values, y_values = remove_adjacent_duplicates(x_values, y_values)

        axes[row_idx, col_idx].plot(x_values, y_values, label=title)
        axes[row_idx, col_idx].set_xlabel(x_axis_label)
        axes[row_idx, col_idx].set_ylabel(y_axis_label)

        if base_line:
            axes[row_idx, col_idx].axhline(
                y=base_line, color="red", linestyle="--", label=base_line_label
            )

        if secondary_base_line:
            axes[row_idx, col_idx].axhline(
                y=secondary_base_line,
                color="blue",
                linestyle="--",
                label=secondary_base_line_label,
            )

        node_idx = idx + 1
        if values["type"] == "Validator":
            node_idx = idx + 1 - trainer_count
        elif values["type"] == "Challenger":
            node_idx = idx + 1 - trainer_count - validator_count

        axes[row_idx, col_idx].legend()
        axes[row_idx, col_idx].set_title(
            f"{label_prefix if label_prefix else values['type']} {node_idx} {subtitle_suffix}"
        )

    plt.tight_layout()
    plt.savefig(output_file)


if __name__ == "__main__":
    trainer_data = get_trainer_token_data("Paper: Scenario 1")
    validator_data = get_validator_token_data("Paper: Scenario 1")
    challenger_data = get_challenger_token_data("Paper: Scenario 1")

    # Trainer Token Graphs
    create_separated_graph(
        data=trainer_data,
        output_file="trainer/scenario-1-seperated.png",
        num_cols=3,
        num_rows=2,
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(12, 8),
        title="Trainer Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=trainer_data,
        output_file="trainer/scenario-1-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(8, 6),
        title="Trainer Token Supply",
    )

    # Validator Token Graphs
    create_separated_graph(
        data=validator_data,
        output_file="validator/scenario-1-seperated.png",
        num_cols=3,
        num_rows=1,
        label_prefix="Validator",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(12, 4),
        title="Validator Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=validator_data,
        output_file="validator/scenario-1-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(8, 6),
        title="Validator Token Supply",
    )

    # Trainer Token Graphs
    create_separated_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-1-seperated.png",
        num_cols=3,
        num_rows=4,
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(12, 8),
        title="Node Token Supply",
        subtitle_suffix="Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
    )
    create_overlayed_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-1-overlayed.png",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(8, 7),
        title="Node Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        secondary_base_line=75000,
        secondary_base_line_label="Initial Token Supply (After Validator Stake)",
    )

    trainer_data = get_trainer_token_data("Paper: Scenario 2")
    validator_data = get_validator_token_data("Paper: Scenario 2")
    challenger_data = get_challenger_token_data("Paper: Scenario 2")

    # Trainer Token Graphs
    create_separated_graph(
        data=trainer_data,
        output_file="trainer/scenario-2-seperated.png",
        num_cols=3,
        num_rows=2,
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(12, 8),
        title="Trainer Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=trainer_data,
        output_file="trainer/scenario-2-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(8, 6),
        title="Trainer Token Supply",
    )

    # Validator Token Graphs
    create_separated_graph(
        data=validator_data,
        output_file="validator/scenario-2-seperated.png",
        num_cols=3,
        num_rows=1,
        label_prefix="Validator",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(12, 4),
        title="Validator Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=validator_data,
        output_file="validator/scenario-2-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(8, 6),
        title="Validator Token Supply",
    )

    # Trainer Token Graphs
    create_separated_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-2-seperated.png",
        num_cols=3,
        num_rows=4,
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(12, 8),
        title="Node Token Supply",
        subtitle_suffix="Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
    )
    create_overlayed_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-2-overlayed.png",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(8, 7),
        title="Node Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        secondary_base_line=75000,
        secondary_base_line_label="Initial Token Supply (After Validator Stake)",
    )

    trainer_data = get_trainer_token_data("Paper: Scenario 3")
    validator_data = get_validator_token_data("Paper: Scenario 3")
    challenger_data = get_challenger_token_data("Paper: Scenario 3")

    # Trainer Token Graphs
    create_separated_graph(
        data=trainer_data,
        output_file="trainer/scenario-3-seperated.png",
        num_cols=3,
        num_rows=2,
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(12, 8),
        title="Trainer Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=trainer_data,
        output_file="trainer/scenario-3-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        fig_size=(8, 6),
        title="Trainer Token Supply",
    )

    # Validator Token Graphs
    create_separated_graph(
        data=validator_data,
        output_file="validator/scenario-3-seperated.png",
        num_cols=3,
        num_rows=1,
        label_prefix="Validator",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(12, 4),
        title="Validator Token Supply",
        subtitle_suffix="Token Supply",
    )
    create_overlayed_graph(
        data=validator_data,
        output_file="validator/scenario-3-overlayed.png",
        label_prefix="Trainer",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        base_line=75000,
        base_line_label="Initial Token Supply (After Stake)",
        fig_size=(8, 6),
        title="Validator Token Supply",
    )

    # Trainer Token Graphs
    create_separated_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-3-seperated.png",
        num_cols=3,
        num_rows=4,
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(12, 8),
        title="Node Token Supply",
        subtitle_suffix="Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
    )
    create_overlayed_graph(
        data={**trainer_data, **validator_data, **challenger_data},
        output_file="combined/scenario-3-overlayed.png",
        x_axis_label="Time (Block Number)",
        y_axis_label="Scatter Token Supply",
        fig_size=(8, 7),
        title="Node Token Supply",
        base_line=100000,
        base_line_label="Initial Token Supply",
        secondary_base_line=75000,
        secondary_base_line_label="Initial Token Supply (After Validator Stake)",
    )

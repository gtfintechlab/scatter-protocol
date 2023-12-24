import os
import sys
from contextlib import contextmanager


@contextmanager
def suppress_stdout():
    with open(os.devnull, "w") as devnull:
        old_stdout = sys.stdout
        sys.stdout = devnull
        try:
            yield
        finally:
            sys.stdout = old_stdout


accuracy = 0
with suppress_stdout():
    import torch
    import torchvision.transforms as transforms
    import torchvision
    from torch.utils.data import DataLoader
    from torchvision.datasets import ImageFolder

    transform = transforms.Compose(
        [
            transforms.Resize((224, 224)),
            transforms.ToTensor(),
            transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
        ]
    )

    num_classes = 2
    eval_dataset = ImageFolder(root="./data/", transform=transform)
    eval_dataloader = DataLoader(eval_dataset, batch_size=64, shuffle=False)

    device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")

    # Load the trained model
    model = torchvision.models.resnet18(weights="ResNet18_Weights.DEFAULT")
    in_features = model.fc.in_features
    model.fc = torch.nn.Linear(in_features, num_classes)
    model.load_state_dict(torch.load("./model.pth"))
    model.to(device)
    model.eval()

    correct_predictions = 0
    total_samples = 0

    with torch.no_grad():
        for images, labels in eval_dataloader:
            images = images.to(device)
            labels = labels.to(device)

            outputs = model(images)
            _, predicted = torch.max(outputs, 1)

            correct_predictions += (predicted == labels).sum().item()
            total_samples += labels.size(0)

    accuracy = correct_predictions / total_samples
print(int(accuracy * 100))

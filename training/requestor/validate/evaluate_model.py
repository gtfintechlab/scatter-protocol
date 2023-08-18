import torch
import torch.nn as nn
import torch.nn.functional as F

import os
import torch
import torch.nn as nn
import torch.optim as optim
import torchvision
import torchvision.transforms as transforms
from torchvision.datasets import ImageFolder
from torch.utils.data import DataLoader
from sklearn.metrics import accuracy_score, precision_score, recall_score
import json


class CustomModel(nn.Module):
    def __init__(
        self,
        root_path,
        num_classes=2,
        batch_size=32,
        num_epochs=10,
        learning_rate=0.001,
    ):
        super(CustomModel, self).__init__()
        self.root_path = root_path
        self.num_classes = num_classes
        self.batch_size = batch_size
        self.num_epochs = num_epochs
        self.learning_rate = learning_rate

        self.transform = transforms.Compose(
            [
                transforms.Resize((224, 224)),
                transforms.ToTensor(),
                transforms.Normalize(
                    mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]
                ),
            ]
        )

        self.dataset = ImageFolder(root=self.root_path, transform=self.transform)
        self.dataloader = DataLoader(
            self.dataset, batch_size=self.batch_size, shuffle=True
        )

        self.model = torchvision.models.resnet18(pretrained=True)
        in_features = self.model.fc.in_features
        self.model.fc = nn.Linear(in_features, self.num_classes)

        self.criterion = nn.CrossEntropyLoss()
        self.optimizer = optim.Adam(self.model.parameters(), lr=self.learning_rate)

    def trainModel(self):
        for epoch in range(self.num_epochs):
            for images, labels in self.dataloader:
                self.optimizer.zero_grad()
                outputs = self.model(images)
                loss = self.criterion(outputs, labels)
                loss.backward()
                self.optimizer.step()
            print(f"Epoch [{epoch+1}/{self.num_epochs}], Loss: {loss.item():.4f}")

    def saveModel(self, save_path):
        torch.save(self.model.state_dict(), save_path)


# Load the trained model
root_path = "./data/"  # Adjust the root path to your dataset location
model = CustomModel(root_path=root_path)
model.load_state_dict(torch.load("model.pth"))
model.eval()

# Initialize lists to store true labels and predicted labels
true_labels = []
predicted_labels = []

# Define a custom evaluation dataset (similar to training dataset)
eval_dataset = ImageFolder(root=root_path, transform=model.transform)
eval_dataloader = DataLoader(eval_dataset, batch_size=model.batch_size, shuffle=False)

# Iterate through the evaluation dataset and make predictions
with torch.no_grad():
    for images, labels in eval_dataloader:
        outputs = model.model(images)  # Access the inner model for inference
        predicted_labels.extend(outputs.argmax(dim=1).cpu().numpy())
        true_labels.extend(labels.numpy())

# Calculate evaluation metrics
accuracy = accuracy_score(true_labels, predicted_labels)
precision = precision_score(true_labels, predicted_labels)
recall = recall_score(true_labels, predicted_labels)

# Create a dictionary to store the metrics
metrics_dict = {"accuracy": accuracy, "precision": precision, "recall": recall}

# Write metrics to a JSON file
with open("output/evaluation_metrics.json", "w+") as json_file:
    json.dump(metrics_dict, json_file)

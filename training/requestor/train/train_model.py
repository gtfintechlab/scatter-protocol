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
dataset = ImageFolder(root="./data/example", transform=transform)
dataloader = DataLoader(dataset, batch_size=64, shuffle=True)

model = torchvision.models.resnet18(weights="ResNet18_Weights.DEFAULT")
in_features = model.fc.in_features
model.fc = torch.nn.Linear(in_features, num_classes)
criterion = torch.nn.CrossEntropyLoss()
optimizer = torch.optim.Adam(model.parameters(), lr=0.001)
numEpochs = 3


def trainModel():
    i = 0
    device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")
    for epoch in range(numEpochs):
        for images, labels in dataloader:
            images.to(device)
            optimizer.zero_grad()
            i += 1
            outputs = model(images)
            loss = criterion(outputs, labels)
            loss.backward()
            optimizer.step()

        print(f"Epoch [{epoch+1}/{numEpochs}], Loss: {loss.item():.4f}")


trainModel()
torch.save(model.state_dict(), "/tmp/output/model.pth")

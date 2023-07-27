import os
import torch
import torch.nn as nn
import torch.nn.functional as F
from training.trainer.data import CustomDataset


class CustomModel(nn.Module):
    def __init__(self, dataset, learningRate=0.01, batchSize=1, epochs=10):
        super(CustomModel, self).__init__()
        # Device to use
        self.device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        self.epochs = epochs
        self.learningRate = learningRate
        self.batchSize = batchSize
        self.epochs = epochs

        # Convolutions
        self.convolutionOne = nn.Conv2d(in_channels=3, out_channels=32, kernel_size=3)
        self.convolutionTwo = nn.Conv2d(in_channels=32, out_channels=64, kernel_size=3)

        # Max Pooling 2D
        self.dropout1 = nn.Dropout2d(0.25)
        self.dropout2 = nn.Dropout2d(0.5)

        # Fully connected Layers
        self.fc1 = nn.Linear(6471936, 128)
        self.fc2 = nn.Linear(128, 2)

        # Data + Data Loaders
        self.trainData = dataset
        self.trainLoader = torch.utils.data.DataLoader(
            dataset=self.trainData, batch_size=self.batchSize, shuffle=True
        )
        self.testLoader = torch.utils.data.DataLoader(
            dataset=self.trainData, batch_size=self.batchSize, shuffle=True
        )

        # Adaptive Learning Rate Method
        self.optimizer = torch.optim.Adadelta(self.parameters(), lr=self.learningRate)

    def forward(self, x):
        # Convolution Layer
        x = self.convolutionOne(x)
        x = F.relu(x)
        x = self.convolutionTwo(x)
        x = F.relu(x)
        # Max Pool 2D
        x = F.max_pool2d(x, 2)
        x = self.dropout1(x)
        x = torch.flatten(x, 1)
        # Fully Connected Layers
        x = self.fc1(x)
        x = F.relu(x)
        x = self.dropout2(x)
        x = self.fc2(x)
        return F.log_softmax(x, dim=1)

    def trainModel(self):
        for epoch in range(1, self.epochs + 1):
            self.train()
            # Iterate through images + labels
            for i, (images, labels) in enumerate(self.trainLoader):
                # Move tensors to the configured device
                images = images.to(self.device)
                labels = labels.to(self.device)

                # Compute Loss
                self.optimizer.zero_grad()
                output = model(images)
                # Compare target to output
                loss = F.nll_loss(output, labels)

                # Backpropagation and optimization
                loss.backward()
                self.optimizer.step()
                if i % 100 == 0:
                    print(
                        "Train Epoch: {} [{}/{} ({:.0f}%)]\tLoss: {:.6f}".format(
                            epoch,
                            i * len(labels),
                            len(self.trainLoader.dataset),
                            100.0 * i / len(self.trainLoader),
                            loss.item(),
                        )
                    )

    def testModel(self):
        self.eval()
        test_loss = 0
        correct = 0
        with torch.no_grad():
            for image, label in self.testLoader:
                image, label = image.to(self.device), label.to(self.device)
                output = model(image)
                # sum up batch loss
                test_loss += F.nll_loss(output, label, reduction="sum").item()
                # get the index of the max probability
                pred = output.argmax(dim=1, keepdim=True)
                correct += pred.eq(label.view_as(pred)).sum().item()
        # Average Loss
        test_loss /= len(self.testLoader.dataset)

        print(
            "\nTest set: Average loss: {:.4f}, Accuracy: {}/{} ({:.0f}%)\n".format(
                test_loss,
                correct,
                len(self.testLoader.dataset),
                100.0 * correct / len(self.testLoader.dataset),
            )
        )

    def saveModel(self, pathName):
        torch.onnx.export(
            model,
            next(iter(model.trainLoader))[0],
            pathName,
            verbose=True,
        )


if __name__ == "__main__":
    dataset = CustomDataset(trainDataPath="../data/example")
    model = CustomModel(dataset=dataset)
    model.eval()
    model.saveModel("models/example.onnx")

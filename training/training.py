# Specify Training Procedure for Your Machine Here

import torch
from data import CustomDataset
import torch.nn.functional as F

model = torch.load('example.pt')
trainData = CustomDataset(trainDataPath='./data/example')
trainLoader = torch.utils.data.DataLoader(dataset=trainData, shuffle=True)
optimizer = torch.optim.Adadelta(model.parameters(), lr=0.01)

running_loss = 0.0
for epoch in range(3):  # loop over the dataset multiple times
    print(f"Loss: {running_loss}", f"Epoch: {epoch}")
    running_loss = 0.0
    for i, data in enumerate(trainLoader, 0):
        # get the inputs; data is a list of [inputs, labels]
        inputs, labels = data

        # zero the parameter gradients
        optimizer.zero_grad()

        # forward + backward + optimize
        outputs = model(inputs)
        loss = F.nll_loss(outputs, labels)
        loss.backward()
        optimizer.step()
        # print statistics
        running_loss += loss.item()

for i, (images, labels) in enumerate(trainLoader):
    sampleInput = images.to("cpu")
    break
        
torch.onnx.export(model, sampleInput, 'example_trained.onnx', verbose=True)
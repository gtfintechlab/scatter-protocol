# Specify Training Procedure for Docker Environment Here

# Run Python 3.7
FROM python:3.7

# Copy this entire folder into the training directory
COPY . /training
WORKDIR /training

# Install dependencies for model
RUN pip3 install -r requirements.txt

# Convert ONNX File to PyTorch File 
RUN python3 convert.py 

# Run training script with data path
RUN python3 training.py $DATA_PATH
# Testing Purposes
# ENTRYPOINT ["tail", "-f", "/dev/null"]
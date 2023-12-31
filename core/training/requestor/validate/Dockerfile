# Specify Evaluation Procedure for Docker Environment Here

# Run Python 3.8
FROM python:3.8

# Copy this entire folder into the training directory
COPY . /tmp

WORKDIR /tmp
RUN echo $(ls)

RUN pip3 install --upgrade pip
RUN pip3 install --default-timeout=100 torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cpu

ENTRYPOINT ["python3", "evaluate_model.py"]

# Testing Purposes
# ENTRYPOINT ["tail", "-f", "/dev/null"]
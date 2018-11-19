# AI Assignment 2
### COMP 472 fall semester 2018

### Step-by-step docker instructions

Install docker and run in the assignment2 directory:

`docker build -t <some-tag-name> .`

Followed by:

`docker run -it -v $(pwd):/assignment2 --rm <some-tag-name>`

### Directory structure

```
|--<Root of project>
|-- ds1
|  |-- info.csv
|  |-- model.pkl
|  |-- test.csv
|  |-- train.csv
|  |-- val.csv
|-- ds2
   |-- info.csv
   |-- model.pkl
   |-- test.csv
   |-- train.csv
   |-- val.csv
```

More specifically, it assume that the `info.csv`, `test.csv`, `train.csv`, and `val.csv` exist in the directory you provide.
Also it expects them to be in a proper format according to what was provided in assignment 2 of COMP 472.

### Running Naive Bayes Bernoulli

`python3 naive-bayes.py (manual | automatic) (inputDirectory) (outputDirectory)`

Where `manual` creates a classfier and output it, and `automatic` takes a classifier as an input.

Example:

`python3 naive-bayes.py automatic ../ds1 ./models/ds1`

Run the above from the `naive-bayes-bernoulli` directory, it will use the `model.pkl` in the `./models/ds1` directory.

### Running Decision Tree

`python decision-tree.py (manual | automatic) (directory)`

Where `manual` creates a classfier and output it, and `automatic` takes a classifier as an input.

Example:

`python decision-tree.py automatic ../ds1`

Run the above from the `decision-tree` directory, it will use the `model.pkl` in the `../ds1` directory.

### Running SVC = 3

`python svc.py (manual | automatic) (directory)`

Where `manual` creates a classfier and output it, and `automatic` takes a classifier as an input.

Example:

`python svc.py automatic ../ds1`

Run the above from the `svc` directory, it will use the `model.pkl` in the `../ds1` directory.

### Running MLPC = 4

`python3 mlpc.py (manual | automatic) (inputDirectory) (outputDirectory)`

Where `manual` creates a classfier and output it, and `automatic` takes a classifier as an input.

Example:

`python3 mlpc.py automatic ../ds1 ./models/ds1`

Run the above from the `mlpc` directory, it will use the `model.pkl` in the `./models/ds1` directory.

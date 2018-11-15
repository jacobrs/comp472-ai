from sklearn.metrics import accuracy_score
from sklearn import tree
from os import sys
import pickle

def getData(filePath, hasLabels=True):
    with open(filePath, "r") as file:
        data = [line.split(",") for line in file.read().split('\n')[:-1]]

    data = [[int(element) for element in row] for row in data]

    if (hasLabels):
        features = [d[:-1] for d in data]
        labels = [d[-1] for d in data]
    else:
        features = data
        labels = None

    return (features, labels)

def treeClassifyWithData(filePath):
    (features, labels) = getData(filePath, True)

    classifier = tree.DecisionTreeClassifier(criterion="entropy")
    classifier.fit(features, labels)

    return classifier

def validateAgainst(classifier, validationFilePath, hasLabels=False):
    (validationFeatures, validationLabels) = getData(validationFilePath, hasLabels)

    predicted = classifier.predict(validationFeatures)

    if (hasLabels):
        accuracy = accuracy_score(validationLabels, predicted)
        print("Validation accuracy: %f" % accuracy)
    
    return predicted

def outputResults(directory, predictedResults, bitmapArray, type):
    with open(directory + "/info.csv", 'r') as file:
        info = [line.split(',') for line in file.read().split('\n')[1:]]
    outputTest = list(map(lambda predicted: info[predicted][1], predictedResults))
    with open(directory + "/" + type + "-dt.csv", 'w') as file:
        for x in range(len(predictedResults)):
            file.write(str(x + 1) + ", " + outputTest[x] + "\n")

directory=sys.argv[2].replace('/', '')

if (sys.argv[1] == "manual"):
    classifier = treeClassifyWithData(directory + "/train.csv")
else:
    # Load the classifier
    with open(directory + "/model.pkl", "rb") as file:
        classifier = pickle.load(file)

print("Using validation set")
validationPrediction = validateAgainst(classifier, directory + "/val.csv", True)
print(validationPrediction)

print("Using test set")
testPrediction = validateAgainst(classifier, directory + "/test.csv")
print(testPrediction)

(testFeatures, testLabels) = getData(directory + "/test.csv", False)
(validationFeatures, validationLabels) = getData(directory + "/val.csv")

# Output mapped predictions for test and validation set
outputResults(directory, testPrediction, testFeatures, "test")
outputResults(directory, validationPrediction, validationFeatures, "val")

if (sys.argv[1] == "manual"):
    # Save the classifier
    with open(directory + "/model.pkl", "wb") as file:
        pickle.dump(classifier, file)
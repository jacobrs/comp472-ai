from sklearn.metrics import accuracy_score, recall_score, precision_score, f1_score, confusion_matrix
from sklearn.naive_bayes import BernoulliNB
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

def naiveBayesBernoulli(filePath):
    (features, labels) = getData(filePath, True)

    classifier = BernoulliNB(alpha=0.5)
    classifier.fit(features, labels)

    return classifier

def validateAgainst(classifier, validationFilePath, outputDirectory, hasLabels=False):
    (validationFeatures, validationLabels) = getData(validationFilePath, hasLabels)

    predicted = classifier.predict(validationFeatures)

    if (hasLabels):
        accuracy = accuracy_score(predicted, validationLabels)
        print("Validation accuracy: %f" % accuracy)
        precision = precision_score(predicted, validationLabels, average='weighted')
        print("Validation precision: %f" % precision)
        recall = recall_score(predicted, validationLabels, average='weighted')
        print("Validation recall: %f" % recall)
        fscore = f1_score(predicted, validationLabels, average='weighted')
        print("Validation fscore: %f" % fscore)

        with open(outputDirectory + "/scores.txt", "w") as file:
            file.write("Accuracy: %f\n" % accuracy)
            file.write("Precision: %f\n" % precision)
            file.write("Recall: %f\n" % recall)
            file.write("fscore: %f\n" % fscore)
        
        with open(outputDirectory + "/confusion_matrix.txt", "w") as file:
            matrix = confusion_matrix(predicted, validationLabels)
            for row in matrix:
                for col in row:
                    file.write(str(col) + " ")
                file.write("\n")
    
    return predicted

def outputResults(directory, predictedResults, typeName):
    outputName = directory.split('/')[-1]
    with open(directory + "/" + outputName + typeName + "-nb.csv", 'w') as file:
        for x in range(len(predictedResults)):
            file.write(str(x + 1) + ", " + str(predictedResults[x]) + "\n")

inputDirectory=sys.argv[2]
outputDirectory=sys.argv[3]

if (sys.argv[1] == "manual"):
    classifier = naiveBayesBernoulli(inputDirectory + "/train.csv")
else:
    # Load the classifier
    with open(outputDirectory + "/model.pkl", "rb") as file:
        classifier = pickle.load(file)

print("Using validation set")
validationPrediction = validateAgainst(classifier, inputDirectory + "/val.csv", outputDirectory, True)

print("Using test set")
testPrediction = validateAgainst(classifier, inputDirectory + "/test.csv", outputDirectory)

# Output mapped predictions for test and validation set
outputResults(outputDirectory, testPrediction, "Test")
outputResults(outputDirectory, validationPrediction, "Val")

if (sys.argv[1] == "manual"):
    # Save the classifier
    with open(outputDirectory + "/model.pkl", "wb") as file:
        pickle.dump(classifier, file)

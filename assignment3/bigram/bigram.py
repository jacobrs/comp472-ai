import os
import re
import math

languages = ['en', 'fr', 'es']
languageMap = {
    'en': 'ENGLISH',
    'fr': 'FRENCH',
    'es': 'SPANISH'
}
countMatrices = {}
chanceMatrices = {}
characterSetRegex = re.compile('[a-zA-Z]')

def createBigramMatrix():
    matrix = [[0.5 for i in range(27)] for j in range(27)]
    return matrix

def appendToBigramMatrix(matrix, openedBook):
    firstChar = True
    for line in openedBook:
        for char in line.strip():
            if characterSetRegex.match(char):
                if firstChar:
                    char = char.lower()
                    prevCharVal = ord(char) - ord('a')
                    firstChar = False
                else:
                    char = char.lower()
                    currCharVal = ord(char) - ord('a')
                    matrix[prevCharVal][currCharVal] += 1
                    prevCharVal = currCharVal

def calculateChanceBigramMatrix(matrix, countMatrix):
    for rowIndex in range(len(countMatrix)):
        totalCount = sum(countMatrix[rowIndex])
        
        for charIndex in range(len(matrix[rowIndex])):
            percentage = countMatrix[rowIndex][charIndex] / totalCount
            matrix[rowIndex][charIndex] = (percentage, math.log(percentage, 10))

def outputMostLikelyLanguage(sentenceNumber, sentence, chanceMatrices, languages):
    with open("../output/out" + str(sentenceNumber) + ".txt", "a") as output:
        output.write('-------------\n')
        output.write('BIGRAM MODEL:\n')
        output.write('\n')

        logTotals = {}
        for lang in languages:
            logTotals[lang] = 0
        firstChar = True
        # Output cumulative log chances
        for char in sentence:
            if characterSetRegex.match(char):
                if firstChar:
                    char = char.lower()
                    prevChar = char
                    prevCharVal = ord(prevChar) - ord('a')
                    firstChar = False
                else:
                    char = char.lower()
                    currCharVal = ord(char) - ord('a')

                    output.write('BIGRAM: %s\n' % (prevChar + char))
                    for lang in languages:
                        percentage, log = chanceMatrices[lang][prevCharVal][currCharVal]
                        logTotals[lang] += log
                        output.write('%s: P(%s|%s) = %f ==> log prob of sentence right now: %f\n' % (lang, char, prevChar, percentage, logTotals[lang]))
                    output.write('\n')

                    prevChar = char
                    prevCharVal = currCharVal
            
        maxVal = None
        maxLang = ''
        for lang in languages:
            if maxVal == None or logTotals[lang] > maxVal:
                maxVal = logTotals[lang]
                maxLang = lang
        
        output.write('According to bigram model, most likely language is %s\n' % languageMap[maxLang])
        output.write('\n')

def outputPercentagesModel(fileLocation):
    pass

# Count bigram letter counts
for lang in languages:
    countMatrices[lang] = createBigramMatrix()

    for file in os.listdir('../books/' + lang):
        if file.endswith('.txt'):
            with open('../books/' + lang + '/' + file) as book:
                appendToBigramMatrix(countMatrices[lang], book)

# Calculate the percentage and log matrices
for lang in languages:
    chanceMatrices[lang] = createBigramMatrix()
    calculateChanceBigramMatrix(chanceMatrices[lang], countMatrices[lang])

# Check most likely language given sentences and output results
with open('../train/sentences.txt') as sentences:
    index = 0
    for line in sentences:
        outputMostLikelyLanguage(index, line.strip(), chanceMatrices, languages)
        index += 1

# Output percentage matrix as a text model
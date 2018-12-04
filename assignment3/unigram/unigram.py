import os
import re
import math

languages = ['en', 'fr', 'es']
languageMap = {
    'en': 'ENGLISH',
    'fr': 'FRENCH',
    'es': 'SPANISH'
}
countDict = {}
chanceDict = {}
characterSetRegex = re.compile('[a-zA-Z]')

delta = 0.5

def appendToUnigramCount(arr, openedBook):
  for line in openedBook:
    for char in line.strip():
      char = chr(char)
      if characterSetRegex.match(char):
        char = char.lower()
        if char in arr :
          arr[char] += 1
        else :
          arr[char] = 1

def calculateChanceUnigramDict(chanceArr, countArr):
  totalCount = sum(countArr.values())
  for k,v in countArr.items():
    percentage = v/totalCount
    chanceArr[k] = (percentage, math.log10(percentage))
    
def outputMostLikelyLanguage(sentenceNumber, sentence, chanceMatrices, languages):
  with open("../output/out" + str(sentenceNumber+ 1) + ".txt", "w") as output:
    output.write('-------------\n')
    output.write('UNIGRAM MODEL:\n')
    output.write('\n')

    logTotals = {}
    for lang in languages:
      logTotals[lang] = 0
    # Output cumulative log chances
    for char in sentence:
      char = char.lower()
      if characterSetRegex.match(char):
        output.write('UNIGRAM: %s\n' % char)
        for lang in languages:
          percentage, log = chanceMatrices[lang][char]
          logTotals[lang] += log
          output.write('%s: P(%s) = %f ==> log prob of sentence right now: %f\n' % (lang, char, percentage, logTotals[lang]))
        output.write('\n')            
    
    maxVal = None
    maxLang = ''
    for lang in languages:
      if maxVal == None or logTotals[lang] > maxVal:
        maxVal = logTotals[lang]
        maxLang = lang
    
    output.write('According to unigram model, most likely language is %s\n' % languageMap[maxLang])
    print("%d: %s" % (sentenceNumber + 1, languageMap[maxLang]))
    output.write('\n')

def outputPercentagesModel(fileLocation):
  pass

# Count bigram letter counts
for lang in languages:
  countDict[lang] = {}
  for i in range(ord('a'), ord('z')+1):
    countDict[lang][chr(i)] = delta

  for file in os.listdir('../books/' + lang):
    if file.endswith('.txt'):
      with open('../books/' + lang + '/' + file, "rb") as book:
        appendToUnigramCount(countDict[lang], book)

# Calculate the percentage and log matrices
for lang in languages:
  chanceDict[lang] = {}
  calculateChanceUnigramDict(chanceDict[lang], countDict[lang])
  langUpper = lang.upper()
  if lang == 'es':
    langUpper = 'OT'
  with open("../output/models/unigram" + langUpper + ".txt", "w") as output:
    for i in range(ord('a'), ord('z')+1):
      output.write("P(%s) = %f\n" % (chr(i), chanceDict[lang][chr(i)][0]))

# Check most likely language given sentences and output results
with open('../train/sentences_unigram.txt') as sentences:
  index = 0
  for line in sentences:
    outputMostLikelyLanguage(index, line.strip(), chanceDict, languages)
    index += 1

# Output percentage matrix as a text model

package main

func findOrder(dependencies [][]rune) []rune {
  adjMap := make(map[rune][]rune)
  visitedList := make([]rune, 0)
  depMap := make(map[rune]int)
  allNodes := make(map[rune]int)
  retVal := make([]rune,0)
	// Replace this placeholder return statement with your code
  for _, set := range dependencies {
      dst := set[0]
      src := set[1]
      if _, ok := adjMap[src]; !ok {
        adjMap[src] = []rune{dst}
      } else {
        adjMap[src] = append(adjMap[src], dst)
      }
      if count, ok := depMap[dst]; !ok {
        depMap[dst] = 1
      } else {
        depMap[dst] = count+1
      }

      if _, ok := depMap[src]; !ok {
        depMap[src] = 0
      }
      allNodes[src] = 0
      allNodes[dst] = 0
  }

  for  node, count := range depMap {
      if count == 0 {
          visitedList = append(visitedList, node)
      }
  }
  
  for {
      if len(visitedList) == 0 {
          break
      }
      currentNode := visitedList[0]
      visitedList = visitedList[1:]
      retVal = append(retVal, currentNode)
    
      for _, adjNode := range adjMap[currentNode] {
          if count, ok := depMap[adjNode]; ok {
            depMap[adjNode] = count - 1
            if count == 1 {
               visitedList = append(visitedList, adjNode)
            }
         }
      }
  }
  if len(retVal) != len(allNodes) {
      retVal = []rune{}
  }

  return retVal
}

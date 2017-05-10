package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
  "strconv"
  "reflect"
)

func main()  {
  texts := readFile("maze3.txt")

  maze := [][]string {}
  for _, v := range texts {
    maze = append(maze, strings.Split(v, ""))
  }

  start := serchWord(maze, "S")
  goal := serchWord(maze, "G")
  goalCount := 0
  count := 1
  strCount := strconv.Itoa(count)
  spaceCount := 0

  for _, v := range nextPoint(start[0]) {
    if maze[v[0]][v[1]] == " " {
      maze[v[0]][v[1]] = strCount
    }
  }

  for goalCount == 0 {
    spaceCount = 0
    for _, point := range serchWord(maze, strCount) {
      strCount = strconv.Itoa(count+1)
      for _, v := range nextPoint(point) {
        if reflect.DeepEqual(v, goal[0]) {
          goalCount = count
        }
        if maze[v[0]][v[1]] == " " {
          maze[v[0]][v[1]] = strCount
          spaceCount++
        }
      }
    }
    count++
    strCount = strconv.Itoa(count)
    if spaceCount == 0 {
      break
    }
  }

  if goalCount == 0 {
    fmt.Println("Fail")
  }else{
    fmt.Println(goalCount)
  }

}

func serchWord (maze [][]string, word string) [][]int {
  list := [][]int {}
  for x, line := range maze {
    for y, char := range line {
      if char == word {
        list = append(list, []int{x, y})
      }
    }
  }
  return list
}


func nextPoint (point []int) [][]int {
  x := point[0]
  y := point[1]
  return [][]int{ {x-1, y}, {x, y+1}, {x+1, y}, {x, y-1} }
}


func readFile(filePath string) []string {
  file, err := os.Open(filePath)
  if err != nil {
    fmt.Print("error")
  }

  texts := []string {}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    texts = append(texts, scanner.Text())
  }

  return texts
}

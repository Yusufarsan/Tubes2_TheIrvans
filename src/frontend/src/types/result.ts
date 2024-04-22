export interface Result {
  result: string[][]
  time_elapsed: number
}

export interface Node {
  id: string
  label: string
}

export interface Edge {
  source: string
  target: string
  id: string
  label: string
}

export interface GraphData {
  nodes: Node[]
  edges: Edge[]
}

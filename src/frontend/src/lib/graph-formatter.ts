import { Edge, Node, Result } from "@/types/result"

export function getTitle(url: string) {
  return url.split("/").pop()?.replace(/_/g, " ")
}

export function graphiphy(data: Result) {
  const uniqueNodes = new Set<string>()

  data.result.forEach((path) => {
    path.forEach((node) => {
      uniqueNodes.add(node)
    })
  })

  const nodes: Node[] = Array.from(uniqueNodes).map((node, index) => {
    if (
      node === data.result[0][0] ||
      node === data.result[0][data.result[0].length - 1]
    ) {
      return {
        id: (index + 1).toString(),
        label: getTitle(node) || "",
        fill: "#25C9F2",
      }
    }

    return {
      id: (index + 1).toString(),
      label: getTitle(node) || "",
    }
  })

  const edges: Edge[] = []

  for (let i = 0; i < data.result.length; i++) {
    const path = data.result[i]

    for (let j = 0; j < path.length - 1; j++) {
      const source = nodes.find((node) => node.label === getTitle(path[j]))
      const target = nodes.find((node) => node.label === getTitle(path[j + 1]))

      if (source && target) {
        edges.push({
          source: source.id,
          target: target.id,
          id: `${source.id}-${target.id}`,
          label: `${source.label}-${target.label}`,
        })
      }
    }
  }

  return { nodes, edges }
}

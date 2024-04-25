import { graphiphy } from "@/lib/graph-formatter"
import { createLazyFileRoute } from "@tanstack/react-router"
import { GraphCanvas } from "reagraph"

export const Route = createLazyFileRoute("/test")({
  component: Test,
})

function Test() {
  const dummy = {
    result: [
      [
        "https://en.wikipedia.org/wiki/Joko_Widodo",
        "https://en.wikipedia.org/wiki/Gadjah_Mada_University",
        "https://en.wikipedia.org/wiki/Philosophy",
      ],
      [
        "https://en.wikipedia.org/wiki/Joko_Widodo",
        "https://en.wikipedia.org/wiki/Pancasila_(politics)",
        "https://en.wikipedia.org/wiki/Philosophy",
      ],
      [
        "https://en.wikipedia.org/wiki/Joko_Widodo",
        "https://en.wikipedia.org/wiki/Sri_Lanka",
        "https://en.wikipedia.org/wiki/Philosophy",
      ],
    ],
    time_elapsed: 8414,
  }

  const { nodes, edges } = graphiphy(dummy)
  return (
    <div className="fixed size-3/4 border border-black">
      <GraphCanvas nodes={nodes} edges={edges} />
    </div>
  )
}

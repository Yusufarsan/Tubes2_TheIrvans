import { createLazyFileRoute } from "@tanstack/react-router"
import HomeSwitch from "@/components/ui/home-switch"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Button } from "@/components/ui/button"
import AutoComplete from "@/components/ui/autocomplete"
import { useState } from "react"
import { GraphData, Result } from "@/types/result"
import { getTitle, graphiphy } from "@/lib/graph-formatter"
import { GraphCanvas, lightTheme } from "reagraph"

export const Route = createLazyFileRoute("/")({
  component: Index,
})

function HomeSelect({
  setSearchMethod,
}: {
  setSearchMethod: (data: string) => void
}) {
  return (
    <Select defaultValue="ids" onValueChange={setSearchMethod}>
      <SelectTrigger className="w-[561px] border-[3px] border-accent bg-foreground/[40%] py-[32px] font-Akaya text-[30px] text-foreground">
        <SelectValue placeholder="Iterative Deepening Search" />
      </SelectTrigger>
      <SelectContent className="w-[561px] bg-foreground font-Akaya text-background">
        <SelectItem className="text-[30px]" value="ids">
          Iterative Deepening Search
        </SelectItem>
        <SelectItem className="text-[30px]" value="bfs">
          Breadth Depth Search
        </SelectItem>
      </SelectContent>
    </Select>
  )
}

const myTheme = {
  ...lightTheme,
  canvas: {
    background: "#DFFDDB",
    fog: "#DFFDDB",
  },
  node: {
    ...lightTheme.node,
    fill: "#061801",
    label: {
      ...lightTheme.node.label,
      color: "#061801",
    },
  },
  edge: {
    ...lightTheme.edge,
    fill: "#0A9B90",
  },
  arrow: {
    ...lightTheme.arrow,
    fill: "#0A9B90",
  },
}

function Index() {
  const [startURL, setStartURL] = useState<string>("")
  const [goalURL, setGoalURL] = useState<string>("")
  const [searchMethod, setSearchMethod] = useState<string>("ids")
  const [showMultiplePath, setShowMultiplePath] = useState<boolean>(false)
  const [result, setResult] = useState<Result | null>(null)
  const [graphData, setGraphData] = useState<GraphData | null>(null)

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    setResult(null)
    setGraphData(null)

    if (showMultiplePath) {
      const data = await fetch(
        `http://localhost:8080/multiple/${searchMethod}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            start: startURL,
            end: goalURL,
          }),
        },
      )

      if (!data.ok) {
        console.error("Failed to fetch data")
        return
      }

      const response = await data.json()
      setResult(response)
      const { nodes, edges } = graphiphy(response)
      setGraphData({ nodes, edges })
    } else {
      const data = await fetch(`http://localhost:8080/single/${searchMethod}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          start: startURL,
          end: goalURL,
        }),
      })

      if (!data.ok) {
        console.error("Failed to fetch data")
        return
      }

      const response = await data.json()
      setResult(response)
      const { nodes, edges } = graphiphy(response)
      setGraphData({ nodes, edges })
    }
  }

  return (
    <form
      method="post"
      onSubmit={handleSubmit}
      className="flex min-h-[calc(100vh-80px)] flex-col items-center gap-[29px] bg-background"
    >
      <p className="mt-[48px] flex justify-center font-Akaya text-[30px] text-foreground">
        Find the shortest path from
      </p>

      <div className="flex w-fit flex-col items-center gap-[29px]">
        <div className="flex items-center gap-[43px]">
          <AutoComplete placeholder="Start Title" setURL={setStartURL} />
          <p className="font-Akaya text-[30px] text-foreground">to</p>
          <AutoComplete placeholder="Goal Title" setURL={setGoalURL} />
        </div>

        <div className=" flex items-center justify-center self-end">
          <p className="mr-[58px] font-Akaya text-[30px] text-foreground">
            Select searching method:
          </p>
          <HomeSelect setSearchMethod={setSearchMethod} />
        </div>
      </div>

      <div className="mr-[180px] flex items-center">
        <p className="mr-[58px] font-Akaya text-[30px] text-foreground">
          Show more than one path:
        </p>
        <HomeSwitch setValue={setShowMultiplePath} />
      </div>

      <div className=" flex items-center justify-center">
        <Button
          className="w-[240px] border-[3px] border-foreground bg-secondary py-[32px] font-Akaya text-[30px] text-foreground hover:bg-accent"
          variant="default"
          size="lg"
          type="submit"
        >
          Find!
        </Button>
      </div>
      {result && graphData && (
        <section className="relative flex min-h-screen flex-col items-center gap-[29px] bg-background">
          <div className="relative">
            <p className="max-w-[600px] text-center font-Akaya text-[30px] text-foreground">
              Found path from{" "}
              <span className="text-accent underline">
                {getTitle(result.result[0][0])}
              </span>{" "}
              to{" "}
              <span className="text-accent underline">
                {getTitle(result.result[0][result.result[0].length - 1])}
              </span>{" "}
              with {graphData.nodes.length} article(s) and{" "}
              {graphData.edges.length} path(s) in {result.time_elapsed / 1000}{" "}
              seconds
            </p>
            <div className="absolute bottom-1/2 left-1/2 top-[350px] h-[400px] w-[800px] -translate-x-1/2 -translate-y-1/2 rounded-md border-[3px] border-accent">
              <GraphCanvas
                nodes={graphData.nodes}
                edges={graphData.edges}
                theme={myTheme}
                layoutType="forceatlas2"
              />
            </div>
          </div>
        </section>
      )}
    </form>
  )
}

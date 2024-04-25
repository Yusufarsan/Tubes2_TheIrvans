import CardPath from "@/components/ui/card-path"
import { createLazyFileRoute } from "@tanstack/react-router"

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

  return (
    <div className="flex max-w-96 flex-col items-center gap-4">
      {dummy.result.map((path, index) => {
        return <CardPath key={index} path={path} />
      })}
    </div>
  )
}

import { getTitle } from "@/lib/graph-formatter"
import * as React from "react"

function CardPath({
  width = 400,
  arr = [
    "https://en.wikipedia.org/wiki/Joko_Widodo",
    "https://en.wikipedia.org/wiki/Vladimir_Putin",
    "https://en.wikipedia.org/wiki/Joe_Biden",
    "https://en.wikipedia.org/wiki/Xi_Jinping",
  ],
  start_end_color = "green-200",
  middle_color = "green-100",
}) {
  return (
    <div className={`w-[${width}px] flex flex-col gap-0 rounded-lg font-Akaya`}>
      {arr.map((path, index) =>
        index === 0 ? (
          <a href={`${path}`} target="_blank">
            <div
              key={index}
              className={`flex items-center gap-2 rounded-t-lg border-x-2 border-t-2 border-background bg-${start_end_color} p-2 transition duration-500 hover:bg-green-300`}
            >
              <div>{getTitle(path)}</div>
            </div>
          </a>
        ) : index === arr.length - 1 ? (
          <a href={`${path}`} target="_blank">
            <div
              key={index}
              className={`flex items-center gap-2 rounded-b-lg border-x-2 border-b-2 border-t border-background bg-${start_end_color} p-2 transition duration-500 hover:bg-green-300`}
            >
              <div>{getTitle(path)}</div>
            </div>
          </a>
        ) : (
          <a href={`${path}`} target="_blank">
            <div
              key={index}
              className={`flex items-center gap-2 border-x-2 border-t border-background bg-${middle_color} p-2 transition duration-500 hover:bg-green-300`}
            >
              <div>{getTitle(path)}</div>
            </div>
          </a>
        ),
      )}
    </div>
  )
}

export default CardPath

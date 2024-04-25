import { getTitle } from "@/lib/graph-formatter"

function CardPath({ path }: { path: string[] }) {
  return (
    <div className="flex h-full min-w-[300px] flex-col rounded-lg border border-background font-Akaya">
      {path.map((node, index) => (
        <a
          href={node}
          target="_blank"
          key={index}
          className={`border-2 border-accent 
          ${index === 0 ? "rounded-t-lg border-b-2" : ""} 
          ${index === path.length - 1 ? "rounded-b-lg border-t-0" : ""} 
          ${index !== 0 && index !== path.length - 1 ? "border-b-2 border-t-0" : ""}
          bg-foreground p-2 text-background transition duration-200 hover:bg-accent`}
        >
          {getTitle(node)}
        </a>
      ))}
    </div>
  )
}

export default CardPath

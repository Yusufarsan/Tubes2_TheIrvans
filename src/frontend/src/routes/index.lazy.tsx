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

export const Route = createLazyFileRoute("/")({
  component: Index,
})

function HomeSelect() {
  return (
    <Select defaultValue="ids">
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

function Index() {
  return (
    <div className="flex flex-col items-center gap-[29px]">
      <p className="mt-[48px] flex justify-center font-Akaya text-[30px] text-foreground">
        Find the shortest path from
      </p>

      <div className="flex w-fit flex-col items-center gap-[29px]">
        <div className="px  flex items-center gap-[43px]">
          <HomeSelect />
          <p className="font-Akaya text-[30px] text-foreground">to</p>
          <HomeSelect />
        </div>

        <div className=" flex items-center justify-center self-end">
          <p className="mr-[58px] font-Akaya text-[30px] text-foreground">
            Select searching method:
          </p>
          <HomeSelect />
        </div>
      </div>

      <div className="mr-[180px] flex items-center">
        <p className="mr-[58px] font-Akaya text-[30px] text-foreground">
          Show more than one path:
        </p>
        <HomeSwitch />
      </div>

      <div className=" flex items-center justify-center">
        <Button
          className="w-[240px] border-[3px] border-foreground bg-secondary py-[32px] font-Akaya text-[30px] text-foreground hover:bg-accent"
          variant="default"
          size="lg"
        >
          Find!
        </Button>
      </div>

      <p className=" max-w-[500px] text-center font-Akaya text-[30px] text-foreground">
        Found path from{" "}
        <span className="text-accent underline">Start Title</span> to{" "}
        <span className="text-accent underline">Goal Title </span>
        with 99 article(s) after checking 999 article(s) in 300.15 seconds
      </p>
    </div>
  )
}

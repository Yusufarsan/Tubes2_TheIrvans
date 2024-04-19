import { createLazyFileRoute } from "@tanstack/react-router"
import InputBox from "@/components/ui/input-box"

export const Route = createLazyFileRoute("/about")({
  component: About,
})

function About() {
  return (
    <>
      <InputBox />
      <div>Hello from about</div>
    </>
  )
}

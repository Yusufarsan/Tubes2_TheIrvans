import { createLazyFileRoute } from "@tanstack/react-router"
import Navbar from "@/components/ui/navbar"

export const Route = createLazyFileRoute("/about")({
  component: About,
})

function About() {
  return (
    <>
      <Navbar />
      {/* <img src="/stima2-bg.jpg" alt="background Image"></img> */}
    </>
  )
}

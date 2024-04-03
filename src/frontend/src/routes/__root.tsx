import { createRootRoute, Link } from "@tanstack/react-router"

export const Route = createRootRoute({
  component: Navbar,
})

function Navbar() {
  return (
    <div className="flex gap-2 p-2">
      <Link to="/" className="[&.active]:font-bold">
        Home
      </Link>
      <Link to="/about" className="[&.active]:font-bold">
        About
      </Link>
    </div>
  )
}

import { createRootRoute, Link, Outlet } from "@tanstack/react-router"
import { ThemeProvider } from "@/components/theme-provider"

const Navbar = () => {
  return (
    <div>
      <div
        className="shadow-inner-2xl  flex flex-row justify-between bg-primary font-Lobster  text-foreground transition-opacity duration-500 ease-in-out"
        data-replace='{ "opacity-0": "opacity-100" }'
      >
        <Link to="/" className="ml-14 p-2">
          <div className="flex flex-row transition hover:scale-90 active:scale-75">
            <img
              className=" flex w-12  scale-90 items-center"
              src="/logo.png"
              alt="logo"
            />
            <div>
              <div className="ml-4 ">
                <div>Tugas Besar 2 Stima</div>
                <div className=" flex font-Akaya  ">By TheIrvans</div>
              </div>
            </div>
          </div>
        </Link>

        <Link
          to="/about"
          className="mr-14 flex transition hover:scale-90 active:scale-75"
        >
          <div className="mr-4  flex items-center">
            <img src="/info.png" alt="info logo" className="shadow-2xl" />
            <div className=" ml-2 shadow-2xl ">About</div>
          </div>
        </Link>
      </div>
    </div>
  )
}

export const Route = createRootRoute({
  component: () => {
    return (
      <>
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
          <Navbar />
          <Outlet />
        </ThemeProvider>
      </>
    )
  },
})

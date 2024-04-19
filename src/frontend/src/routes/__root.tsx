import { createRootRoute, Link, Outlet } from "@tanstack/react-router"
import { ThemeProvider } from "@/components/theme-provider"

const Navbar = () => {
  return (
    <nav>
      <div className="flex h-20 items-center justify-between bg-primary px-14 font-Lobster text-foreground shadow-md shadow-slate-500">
        <Link to="/">
          <div className="flex items-center gap-4 transition hover:scale-90 active:scale-75">
            <img
              className="flex w-12 scale-90 items-center shadow-none"
              src="/logo.png"
              alt="logo"
            />
            <div className="flex flex-col">
              <p className="text-[28px] shadow-slate-500">
                Tugas Besar 2 Stima
              </p>
              <p className="font-Akaya text-[15px]">By TheIrvans</p>
            </div>
          </div>
        </Link>

        <Link
          to="/about"
          className=" flex transition hover:scale-90 active:scale-75"
        >
          <div className="flex items-center gap-4">
            <img
              src="/info.png"
              alt="info logo"
              className="scale-25 drop-shadow-2xl"
            />
            <div className="text-[24px] shadow-2xl">About</div>
          </div>
        </Link>
      </div>
    </nav>
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

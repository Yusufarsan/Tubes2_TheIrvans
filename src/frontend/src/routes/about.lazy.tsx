import { createLazyFileRoute } from "@tanstack/react-router"
import { Separator } from "@/components/ui/separator"

export const Route = createLazyFileRoute("/about")({
  component: About,
})

function About() {
  return (
    <>
      <main className=" relative flex min-h-[calc(100vh-80px)] flex-col gap-12 bg-background text-justify font-Akaya text-[40px] text-foreground">
        <img
          src="stima2-bg.jpg"
          alt="Background Image"
          className="absolute h-full w-full object-cover"
        />
        <div className="absolute h-full w-full bg-background opacity-80"></div>
        <section className="relative px-20 ">
          <h1 className="py-5 text-center">About This Project</h1>
          <div className="flex flex-col gap-4  text-[20px]">
            <p>
              WikiRace or Wiki Game is a game that involves Wikipedia, a free
              online encyclopedia managed by various volunteers around the
              world, where players start at a Wikipedia article and must
              navigate through other articles on Wikipedia (by clicking links
              within each article) to reach a predetermined article in the
              shortest time or with the fewest clicks (articles)
            </p>
            <p>
              This website aims to find the shortest article route from one
              article to another. After you enter the initial article title and
              the destination article title, and click the “Find!” button, it
              will display:
              <ol className="list-disc pl-8">
                <li>The number of articles checked</li>
                <li>The number of articles passed through</li>
                <li>The route of article exploration</li>
                <li>The search time, and</li>
                <li>Graph visualization of the route exploration</li>
              </ol>
            </p>
            <p>
              You can choose the search method with IDS (Iterative Deepening
              Search) or BFS (Breadth First Search). IDS will search by
              increasing the depth-cutoff value using a series of DFS (Depth
              First Search) until a solution is found. In searching for a node
              in a graph, DFS will search by expanding the first root child of
              the chosen search tree and go deeper and deeper until the target
              node is found, or until it finds a node that has no children.
              Meanwhile, BFS will start the graph search from the root node and
              then explore all its neighboring nodes. Then, for each of these
              closest nodes, it explores neighboring nodes that have not been
              checked, and so on, until the target node is found.
            </p>
            <p>
              In addition, you can also set constraint in searching for the
              article exploration route so that the shortest route displayed is
              not only one, but all the shortest routes from the search results
              will also be displayed.
            </p>
          </div>
        </section>

        <Separator className="relative max-w-[1500px] self-center bg-white" />
        <section className="relative -bottom-12 flex flex-col  gap-4  text-center">
          <h1>Meet Our Team</h1>
          <div className="flex justify-center gap-[80px] text-[30px] ">
            {/* Card Shafiq */}
            <div className="relative flex flex-col">
              <div className="absolute bottom-0 h-1/2 w-full bg-gradient-to-t from-background to-transparent "></div>
              <img
                src="safik.png"
                alt=""
                className="w-[250px] rounded-t-[85px] "
              />
              <div className="absolute top-[200px] pl-2 text-start">
                <p className="text">Shafiq Irvansyah</p>
                <p>13522003</p>
                <p className="text-[20px]">"Papalepale"</p>
              </div>
            </div>
            {/* Card bacin */}
            <div className="relative bottom-0 text-start">
              <div className="absolute bottom-0 h-1/2 w-full bg-gradient-to-t from-background to-transparent "></div>
              <img
                src="bacin.png"
                alt=""
                className="w-[250px]  rounded-t-[85px] "
              />
              <div className="absolute top-[200px] pl-2 text-start">
                <p>Ahmad Naufal R.</p>
                <p>13522005</p>
                <p className="text-[20px]">"quote"</p>
              </div>
            </div>
            {/* Card ucup */}
            <div className="relative bottom-0 text-start">
              <div className="absolute bottom-0 h-1/2 w-full bg-gradient-to-t from-background to-transparent "></div>
              <img
                src="ucup.png"
                alt=""
                className="w-[250px]  rounded-t-[85px] "
              />
              <div className="absolute top-[200px] pl-2 text-start">
                <p>Yusuf Ardian Sandi</p>
                <p>135220015</p>
                <p className="text-[20px]">"Asisten candu web dev ya, Hehe"</p>
              </div>
            </div>
          </div>
        </section>
      </main>
      <div className="h-[200px] bg-background"></div>
    </>
  )
}

const Navbar = () => {
  return (
    <div>
      <div
        className="shadow-inner-2xl before:opa flex flex-row justify-between bg-primary p-2 font-Lobster text-text transition-opacity duration-500 ease-in-out"
        data-replace='{ "opacity-0": "opacity-100" }'
      >
        <div className="flex flex-row">
          <img className=" flex w-12 items-center" src="/logo.png" alt="logo" />
          <div>
            <div className="ml-4">
              <div>Tugas Besar 2 Stima</div>
              <div className=" flex font-Akaya">By TheIrvans</div>
            </div>
          </div>
        </div>

        <div className="mr-4 flex items-center ">
          {" "}
          <img src="/info.png" alt="info logo" className="shadow-2xl" />
          <div className=" ml-2 shadow-2xl ">About</div>
        </div>
      </div>
    </div>
  )
}

export default Navbar

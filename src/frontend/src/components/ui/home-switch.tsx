"use client"

import { Switch } from "@/components/ui/switch"
import { useState } from "react"

function HomeSwitch() {
  const [checked, setChecked] = useState(false)
  console.log(checked)

  return (
    <>
      <div className="flex space-x-1 font-Akaya items-center">
        <Switch onClick={() => setChecked(!checked)} />
        {checked ? <p>On</p> : <p>Off</p>}
      </div>
    </>
  )
}

export default HomeSwitch

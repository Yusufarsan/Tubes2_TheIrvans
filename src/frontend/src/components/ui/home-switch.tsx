"use client"

import { Switch } from "@/components/ui/switch"
import { useState } from "react"

function HomeSwitch() {
  const [checked, setChecked] = useState(false)
  console.log(checked)

  return (
    <>
      <div className="font-Akaya text-[30px] text-foreground flex space-x-[10px] font-Akaya items-center">
        <Switch onClick={() => setChecked(!checked)} />
        {checked ? <p>On</p> : <p>Off</p>}
      </div>
    </>
  )
}

export default HomeSwitch

"use client"

import { Switch } from "@/components/ui/switch"
import { useState } from "react"

function HomeSwitch({ setValue }: { setValue: (val: boolean) => void }) {
  const [checked, setChecked] = useState(false)

  return (
    <>
      <div className="flex items-center space-x-[10px] font-Akaya text-[30px] text-foreground">
        <Switch
          onClick={() => {
            setChecked(!checked)
            setValue(!checked)
          }}
        />
        {checked ? <p>On</p> : <p>Off</p>}
      </div>
    </>
  )
}

export default HomeSwitch

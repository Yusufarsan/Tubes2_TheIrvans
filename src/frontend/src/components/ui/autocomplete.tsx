"use client"

import getWiki from "@/lib/wiki-search-handler"
import { ItemData } from "@/types/wiki"
import { useEffect, useState } from "react"

import {
  Command,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command"

function AutoComplete({
  placeholder,
  setURL,
}: {
  placeholder: string
  setURL: (data: string) => void
}) {
  const [query, setQuery] = useState<{ title: ItemData[]; url: ItemData[] }>({
    title: [],
    url: [],
  })
  const [input, setInput] = useState<string>("")
  const [value, setValue] = useState<string>("")
  const [visible, setVisible] = useState<boolean>(false)

  useEffect(() => {
    const delay = setTimeout(() => {
      getWiki(input).then((data) => {
        setQuery(data)
        setVisible(true)
      })
    }, 500)

    return () => {
      clearTimeout(delay)
    }
  }, [input])

  return (
    <Command className="relative w-[561px] overflow-visible">
      <CommandInput
        value={value}
        onInput={(e) => {
          setInput(e.currentTarget.value)
          setValue(e.currentTarget.value)
        }}
        placeholder={placeholder}
        className="w-full border-[3px] border-accent bg-foreground py-[32px] text-center font-Akaya text-[30px] text-background placeholder:text-center placeholder:text-background"
      />
      {query.title.length === 0 ? (
        <CommandList></CommandList>
      ) : (
        <CommandList
          className={`absolute top-[70px] z-50 w-[561px] rounded-md border-[3px] border-accent bg-foreground ${visible ? "block" : "hidden"}`}
        >
          {query.title.map((title, index) => (
            <CommandItem
              key={index}
              onSelect={() => {
                setValue(title.name)
                setVisible(false)
                setURL(query.url[index].name)
              }}
              className="font-Akaya text-lg text-background"
            >
              {title.name}
            </CommandItem>
          ))}
        </CommandList>
      )}
    </Command>
  )
}

export default AutoComplete

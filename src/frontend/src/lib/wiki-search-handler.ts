import { ItemData } from "@/types/wiki"

async function getWiki(
  name: string,
): Promise<{ title: ItemData[]; url: ItemData[] }> {
  const response = await fetch(
    `https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=${name}&origin=*`,
  )

  const data = await response.json()

  if (data?.error) {
    return { title: [], url: [] }
  }

  const titleData: ItemData[] = data[1].map((title: string, index: number) => {
    return { id: index, name: title }
  })

  const urlData: ItemData[] = data[3].map((url: string, index: number) => {
    return { id: index, name: url }
  })

  return { title: titleData, url: urlData }
}

export default getWiki

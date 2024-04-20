async function getWiki(name: string): Promise<[string[], string[]]> {
  const response = await fetch(
    `https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=${name}&origin=*`,
  )
  const data = await response.json()
  console.log(data)
  return [data[1], data[3]]
}

export default getWiki

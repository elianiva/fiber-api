import { writable, Writable } from "svelte/store"

interface Book {
  _id: string
  name: string
  author: string
  pages: number
  imgUrl: string
}

interface Result {
  status: number
  message: string
  data: Book[]
}

const cache: Map<string, Result> = new Map<any, any>()

/*
 * Function that returns result or cache if the request has been done before
 *
 * @param {string} url - Request URL that needs to be fetched
 * @return {Writable} - Store that contains the result / cache
 */
export const fetchData = (url: string): Writable<Promise<Result>> => {
  const store: Writable<Promise<Result>> = writable(new Promise(() => {}))

  // if `cache` has the item with same `url` key then set the result
  // to the old cache
  if (cache.has(url)) store.set(Promise.resolve(cache.get(url)))
  ;(async () => {
    const res: Response = await fetch(url)
    const data: Result = await res.json()

    // set both cache and store when the result has been found
    cache.set(url, data)
    store.set(Promise.resolve(data))
  })()

  return store
}

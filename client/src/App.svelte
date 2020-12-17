<style>
h1 {
  font-size: 3rem;
  font-weight: 200;
  font-family: "Raleway", sans-serif;
  text-align: center;
  color: #82aaff;
  margin-top: 2rem;
}

.main {
  max-width: 960px;
  margin: 0 auto;
}

.main__book {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(28rem, 1fr));
  gap: 1rem;
  margin-top: 1rem;
}

.main__action {
  margin: 1rem 0 2rem;
}

.action__add {
  border: none;
  outline: none;
  padding: 0.5rem;
  font-size: 1.125rem;
  border-radius: 0.25rem;
  background-color: #34d399;
  color: #ffffff;
  margin: 0 auto;
  display: block;
  cursor: pointer;
}

.action__add:hover {
  filter: brightness(0.95);
}
</style>

<svelte:head>
  <link rel="preconnect" href="https://fonts.gstatic.com" />
  <link
    href="https://fonts.googleapis.com/css2?family=Raleway:wght@200&display=swap"
    rel="stylesheet"
  />
  <link
    href="https://fonts.googleapis.com/css2?family=Open+Sans&display=swap"
    rel="stylesheet"
  />
</svelte:head>

<main class="main">
  <h1>Here be books, etc etc.</h1>
  <div class="main__action">
    <button class="action__add">Add New Book</button>
  </div>
  <div class="main__book">
    {#await $books}
      <span>loading...</span>
    {:then result}
      {#each result.data as { _id, name, author, pages, imgUrl }}
        <Card id={_id} {name} {author} {pages} {imgUrl} onClick={removeData} />
      {/each}
    {/await}
  </div>
</main>

<script lang="ts">
import type { Writable } from "svelte/store"
import { fetchData } from "./helpers/fetch"
import Card from "./components/Card.svelte"

let books: Writable<Promise<{ [key: string]: any }>> = fetchData(
  "http://localhost:3000/api/book"
)

const removeData = (id: string): void => {
  fetch(`http://localhost:3000/api/book/id/${id}`, {
    method: "DELETE",
  })
    .then(res => res.json())
    .then(res => {
      books = fetchData("http://localhost:3000/api/book")

      console.log(res.message)
    })
}
</script>

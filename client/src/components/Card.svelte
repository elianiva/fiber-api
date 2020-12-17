<style>
.card {
  padding: 0.75rem;
  box-shadow: 0 0.25rem 1rem rgba(0, 0, 0, 0.075);
  display: flex;
  gap: 1rem;
  border-radius: 0.25rem;
}

.card__img {
  width: 10rem;
  border-radius: 0.25rem;
}

.card__detail {
  display: flex;
  flex-direction: column;
  color: #454545;
  font-family: "Open Sans", sans-serif;
}

.card__name {
  text-transform: capitalize;
  font-size: 1.5rem;
  line-height: 1.5em;
}

.card__pages,
.card__author {
  font-size: 1.125rem;
  line-height: 1.5em;
}

.card__action {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

.action__update,
.action__remove {
  border: none;
  outline: none;
  padding: 0.5rem;
  font-size: 1.125rem;
  border-radius: 0.25rem;
  cursor: pointer;
}

.action__update:hover,
.action__remove:hover {
  filter: brightness(0.95);
}

.action__update {
  background-color: #60a5fa;
  color: #ffffff;
}

.action__remove {
  background-color: #ef4444;
  color: #ffffff;
}
</style>

<div class="card">
  <img class="card__img" src="http://localhost:3000{imgUrl}" alt="book_cover" />
  <div class="card__detail">
    <span class="card__name">Name: {name}</span>
    <span class="card__author">Author: {author}</span>
    <span class="card__pages">Pages: {pages}</span>
    <div class="card__action">
      <button class="action__update">Update</button>
      <button
        class="action__remove"
        on:click={() => onClick(id)}
      >Remove</button>
    </div>
  </div>
</div>

<script lang="ts">
export let name: string
export let author: string
export let pages: number
export let imgUrl: string
export let id: string
export let onClick: (id: string) => void
</script>

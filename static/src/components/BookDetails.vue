<template>
  <div class="text-gray-500 mt-4">
    <div>
      <b class="font-bold">ISBN: </b> {{ isbn }}
    </div>
    <div>
      <b class="font-bold">Publisher: </b> {{ book.volumeInfo.publisher }}
    </div>
    <div>
      <b class="font-bold">Published Date: </b> {{ book.volumeInfo.publishedDate }}
    </div>
    <div>
      <b class="font-bold">Categories: </b> {{ book.volumeInfo.categories.join(', ').trim(', ') }}
    </div>
    <div>
      <b class="font-bold">Mature: </b> {{ book.volumeInfo.maturityRating === "NOT_MATURE" ? "No" : "Yes" }}
    </div>

    <div>
      <b class="font-bold">Page Count: </b> {{ book.volumeInfo.pageCount }}
    </div>

  </div>
</template>

<script setup>
const props = defineProps({
  isbn: String,
  book:Object
})
</script>

<style scoped>

</style>
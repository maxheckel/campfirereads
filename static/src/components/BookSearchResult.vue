<template>
  <div>
    <a :href="bookHref(book)" class="hidden md:grid grid grid-cols-[20%_80%] my-8 border-b-1 border-b-gray-200 py-4 border-box">

      <img v-if="imageUrl(book)" :src="imageUrl(book)" class="w-2/3  block relative mx-auto border">
      <div v-else class="w-2/3 min-w-2/3 h-[240px] bg-gray-200 relative mx-auto">

      </div>
      <div class="relative">
        <h2 class="text-2xl">{{book.volumeInfo.title}}</h2>
        <div class="italic" v-if="book?.volumeInfo?.authors?.length > 0">By {{book?.volumeInfo?.authors?.join(', ').trim(', ')}}</div>
        <div class="text-sm">Published date {{book.volumeInfo.publishedDate}}</div>
        <div class="max-h-20 overflow-hidden block">
          {{book.volumeInfo.description.split(' ').slice(0, 50).join(' ')}}...
        </div>
      </div>

    </a>
    <Book class="md:hidden" :extended="true" :book="book"/>
  </div>

</template>

<script setup>
import {bookHref, imageUrl} from "../services/utils";
import Book from "../components/Book.vue";
const props = defineProps({
  book: Object
})
</script>

<style scoped>

</style>
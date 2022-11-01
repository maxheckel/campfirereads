<template>
  <div class="w-full relative">

    <Header class="justify-center" :text="props.title" :icon-path="'/media/pine.svg'"/>
    <div v-if="data.loading" class="relative w-full ml-auto text-center">
      <Loading class="relative mx-auto opacity-50"/>
    </div>
    <div v-if="!data.loading" class="grid grid-cols-2 md:grid-cols-6 gap-10">
      <Book v-for="book in data.books.slice(0, 6)" :book="book"></Book>
    </div>
    <div class="flex justify-center">
      <a :href="categoryLink" class="inline-block bg-np-yellow-200 px-4 py-2 rounded-lg text-xl font-light cursor-pointer">View All</a>
    </div>
  </div>

</template>

<script setup>
import {onMounted, reactive} from "vue";
import Loading from "./icons/Loading.vue";
import Book from "./Book.vue";
import Header from "./Header.vue";

const props = defineProps({
  endpoint: String,
  title: String,
  books: Array,
    categoryLink: String
})

const data = reactive({
  books: props.books ? props.books : [],
  loading: true
})

onMounted(() => {
  if( props.books !== undefined && props.books.length > 0){
    data.loading = false;
    return;
  }
  fetch(import.meta.env.VITE_API_HOST + props.endpoint)
      .then((response) => response.json())
      .then((resp) => {
        data.books = resp;
        data.loading = false;
      });
})
</script>

<style scoped>

</style>
<template>
  <a :href="bookHref(book)" class="mt-8 mb-4 cover">
    <div class="flex items-center justify-center center">
      <img v-if="imageUrl() != ''" class="w-full transition-transform" :src="imageUrl()">
      <div v-else class="w-full h-[250px] bg-gray-200 block relative text-4xl overflow-hidden font-bold text-gray-300">{{props.book.volumeInfo.title}}</div>
    </div>
    <div class="w-full text-center px-2 mt-4">
      <span class="font-bold block"> {{props.book.volumeInfo.title}}</span>
      <span v-if="props.book.volumeInfo?.authors?.length > 0">
        By: {{props.book?.volumeInfo?.authors?.join(", ").trim(", ")}}
        </span>
      <span v-if="extended" class="text-sm text-gray-500">
        Publish date:  {{book.volumeInfo.publishedDate}}
      </span>
    </div>
  </a>

</template>

<script setup>
import {bookHref} from "../services/utils.js";
const props = defineProps({
  book: Object,
  extended: Boolean
})

function imageUrl(){
  return props.book?.volumeInfo?.imageLinks?.thumbnail.replace("edge=curl", "").replace("http://", "https://")
}

</script>

<style scoped lang="scss">
  .cover{
    position: relative;
    &:hover{
        cursor:pointer;
        img{
          transform: scale(1.1);

        }
    }
  }

</style>
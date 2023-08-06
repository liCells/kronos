<template>
  <div>
    <div class="inset-x-1/4 top-60 z-50">
      <input
          ref="searchInput"
          type="text"
          class="w-full px-4 py-2 text-lg border border-gray-500 rounded-md outline-none bg-gray-100 dark:md:focus:bg-searchInput"
          v-model="q"
          @keyup.enter="search(false)"
      />
    </div>

    <div>
      <div v-if="res && res.total > 0" v-for="(item, i) in res.data" id="scrollContainer">
        <ul>
          {{ i + 1 }} - <a :href="item._source[mapping[item._index].source]">{{
            item._source[mapping[item._index].title]
          }}</a> -
          {{ item._source[mapping[item._index].tag] }}
        </ul>
        <div class="flex">
          <div class="basis-1/4 text-slate-300 truncate">
            {{ item._source[mapping[item._index].content] }}
          </div>
        </div>
      </div>
    </div>
    <div v-if="showLoadButton()">
      <button @click="loadNext">Load Next</button>
    </div>
  </div>
</template>

<script>
import req from './api/axios'

export default {
  name: 'App',
  data() {
    return {
      page: {
        from: 0,
        size: 30,
      },
      q: '',
      mapping: {},
      res: undefined,
    }
  },
  mounted() {
    this.getMapping()
  },
  methods: {
    async search(append) {
      if (!append) {
        this.page.from = 0
        this.page.size = 30
      }
      req({
        url: '/data/basic/search',
        method: "POST",
        data: {
          content: this.q,
          indexes: [],
          from: this.page.from,
          size: this.page.size,
        },
      }).then(res => {
        if (res.code === 200) {
          if (append) {
            this.res.data.push(...res.data.data)
          } else {
            this.res = res.data
          }
        }
      })
    },
    async getMapping() {
      req({
        url: '/mapping/get',
        method: "GET",
      }).then(res => {
        if (res.code === 200) {
          this.mapping = res.data
        }
      })
    },
    showLoadButton() {
      let max = this.page.from + this.page.size
      return this.res !== undefined && max < this.res.total && this.res.total !== 0
    },
    loadNext() {
      this.page.from += this.page.size
      this.search(true)
    },
  },
}
</script>

<style>
</style>

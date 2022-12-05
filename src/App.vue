<template>
  <h1>Game of life</h1>

  <button @click="setDimensions(10)">10x10</button>
  <button @click="setDimensions(20)">20x20</button>
  <button @click="setDimensions(50)">50x50</button>
  <button @click="setDimensions(100)">100x100</button>

  <form class="options" @submit.prevent="setGrid">
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Generate a seed!</button>
  </form>

  <button @click="playOrStop">{{play_stop}}</button>
  
  <Grid :grid="grid"/>

</template>

<script>
import Grid from './components/Grid.vue'

var state

export default {
  name: 'App',
  components: {
    Grid
  },

  data() {
    return{
      size: 0,
      probability: 0.0,
      grid: [],
      play_stop: "Play!",
    }
  },

  methods:{
    setDimensions(size){
      this.size = size
    },

    setGrid() {
      this.grid = []
      if (this.size == 0 || this.probability == 0.0 ) {
        alert('Please chose any size and/or set the probabiliy.')
      } else {
        for(let i=0; i<this.size; i++){
          this.grid[i] = []
          for (let j = 0; j < this.size; j++) {
            if(this.randomNumber() > (this.probability)*10){
              this.grid[i][j] = false 
            } 
            else{
              this.grid[i][j] = true
            } 
          }
        }
      }
    },

    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
    },

    playOrStop() {
      if (this.size == []) {
        alert('Please chose generate a seed.')
      } else {
        if(this.play_stop == "Play!"){
          this.play_stop = "Stop"
          state = setInterval(this.aplyRules, 100);
        } else {
          this.play_stop = "Play!"
          clearInterval(state)
        }
      }
    },

    aplyRules() {
      for (let i = 0; i < this.grid.length; i++) {
        for (let j = 0; j < this.grid.length; j++) {
          
          let live_neighbors = this.liveNeighbors(j, i)
          
          if ((this.grid[i][j]) && (live_neighbors < 2))
            this.grid[i][j] = false

          else if ((this.grid[i][j]) && (live_neighbors > 3))
            this.grid[i][j] = false

          else if ((!this.grid[i][j]) && (live_neighbors == 3))
            this.grid[i][j] = true
        }
      }
    },

    liveNeighbors(x, y){
      let count = 0
      for (let i = y - 1; i <= y+1; i++) {
        for (let j = x - 1; j <= x+1; j++) {
          if ((i == y && j == x) || (i < 0 || j < 0) || (i >= this.grid.length || j >= this.grid.length) ){
            continue
          }
          if (this.grid[i][j]){
            count++
          }
        }
      }
      return count
    }

  }
}
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 3rem;
  }

  .number-input, button{
    margin: 0.5rem 0.25rem;
  }
</style>
<template>
  <h1>Game of life</h1>

  <button @click="setDimensions(10)">10x10</button>
  <button @click="setDimensions(20)">20x20</button>
  <button @click="setDimensions(50)">50x50</button>
  <button @click="setDimensions(100)">100x100</button>

  <form class="options" @submit.prevent="setGrid">
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Play!</button>
  </form>
  
  <table>
    <tr v-for="(row, rowKey) in cells" :key="rowKey">
      <td v-for="(col, colKey) in row" :key="colKey" 
      :class="col.alive ? 'alive' : 'dead'" class="cell"
      ></td>
    </tr>
  </table>
</template>

<script>
export default {
  name: 'App',
  components: {
  },

  data() {
    return{
      size: 0,
      probability: 0.0,
      cells: []
    }
  },

  methods:{
    setDimensions(size){
      this.size = size
    },
    setGrid() {
      this.cells = []
      for(let i=0; i<this.size; i++){
        this.cells[i] = []
        for (let j = 0; j < this.size; j++) {
          if(this.randomNumber() > (this.probability)*10) this.cells[i][j]={ alive: false }
          else this.cells[i][j]={ alive: true }
        }
      }

    },
    changeState(cell, i,j) {
      this.cells[i][j].alive = !cell.alive
      console.log(cell)
    },
    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
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
  margin: 1rem 0.25rem;
}

.cell {
  padding: 0.26rem;
}

.alive {
  background-color: black !important;
}

.dead {
  border-style:solid;
  background-color: white ;
  border-width: 1px;
}

table {
  margin: auto;
}

</style>
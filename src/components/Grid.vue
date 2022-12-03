<script>
export default {
  props: {
    grid: {
        type: []
    },
  }
}
</script>

<template>
   <table>
      <tr v-for="(row, rowKey) in grid" :key="rowKey">
        <td v-for="(col, colKey) in row" :key="colKey" 
        :class="col ? 'alive' : 'dead'" class="cell"
        ></td>
      </tr>
    </table>
</template>

<style>
  .cell {
    padding: 0.26rem;
  }
  
  .alive {
    border-style:solid;
    background-color: black ;
    border-width: 1px;
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
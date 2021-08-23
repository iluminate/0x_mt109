<template>
  <div>
    <div v-for="row in data.routes" class="maptag-card maptag-bleft-blue">
      <div v-for="col in cols" class="maptag-row">
        <span>{{ col.label }}</span>
        <br/>
        <p :class="col.class">{{ row[col.key] || "-" }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  layout: 'default',
  data() {
    return {
      cols: [
        { label: "Ruta", key: "id", size: 1, class: "" },
        { label: "Conductor", key: "driver_name", size: 3, class: "" },
        { label: "Creada", key: "created_at", size: 2, class: "" },
        { label: "Inicio", key: "started_at", size: 1, class: "" },
        { label: "Terminó", key: "ended_at", size: 1, class: "" },
        { label: "Entregas", key: "deliveries", size: 1, class: "" },
        { label: "", key: "status", size: 1, class: "tag-blue" }
      ],
      data: {
        "routes":[
          {
            "id":1,
            "driver_name":"Ali",
            "created_at":"2020-11-03T20:53:31-06:00",
            "deliveries":0,
            "status":"pending"
          },
          {
            "id":2,
            "driver_name":"Sam",
            "created_at":"2020-12-01T00:09:07-06:00",
            "deliveries":0,
            "status":"pending"
          },
          {
            "id":3,
            "driver_name":"Kiley",
            "created_at":"2020-11-22T21:02:58-06:00",
            "deliveries":0,
            "status":"pending"
          }
        ]
      }
    }
  },
})
</script>


<style>
.maptag-card {
  margin: 6px;
  padding: 6px;
  background-color: #ffffff;
  -webkit-border-radius: 3px;
  -moz-border-radius: 3px;
  border-radius: 3px;
}
.maptag-bleft-blue {
  border-left-style: solid;
  border-left-width: 4px;
  border-left-color: #3b8bff;
}
.maptag-row {
  display: inline-block;
  padding: 8px 22px;
}

.maptag-row span {
  color: #999999;
  font-size: 0.85em;
}
p {
  margin: 0px;
}
.tag-blue {
  padding: 4px 12px;
  color: #e2f2ff;
  background-color: #3b8bff;
  -webkit-border-radius: 20px;
  -moz-border-radius: 20px;
  border-radius: 20px;
  font-size: 0.85em;
}
</style>

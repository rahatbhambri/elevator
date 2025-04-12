<template>
  <div class="container">
    <h2>Elevator Simulation</h2>

    <button v-if="!started" @click="startElevator" class="start-btn">
      Start Elevator
    </button>

    <div v-else class="elevator-shaft">
      <div
        v-for="floor in [...Array(totalFloors).keys()].reverse()"
        :key="floor"
        class="floor"
      >
        <div class="floor-label">
          <!-- Display correct floor numbering, where 0 is at the bottom and 4 is at the top -->
          Floor {{ totalFloors - 1 - floor }}
          <div>
            <!-- For the topmost floor (Floor 4), only Down button -->
            <button
              v-if="totalFloors - 1 - floor === totalFloors - 1"
              @click="pressButton('down', totalFloors - 1 - floor)"
              class="btn"
            >
              ↓
            </button>

            <!-- For the bottommost floor (Floor 0), only Up button -->
            <button
              v-if="totalFloors - 1 - floor === 0"
              @click="pressButton('up', totalFloors - 1 - floor)"
              class="btn"
            >
              ↑
            </button>

            <!-- For intermediate floors, show both Up and Down buttons -->
            <button
              v-if="totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1"
              @click="pressButton('up', totalFloors - 1 - floor)"
              class="btn"
            >
              ↑
            </button>
            <button
              v-if="totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1"
              @click="pressButton('down', totalFloors - 1 - floor)"
              class="btn"
            >
              ↓
            </button>
          </div>
        </div>
      </div>

      <div class="elevator" :style="{ bottom: elevatorPosition + 'px' }"></div>
    </div>

    <audio ref="elevatorSound" src="/grrr.mp3" preload="auto" />
  </div>
</template>

<script setup>
import { ref } from 'vue'

const totalFloors = 5
const floorHeight = 100
const currentFloor = ref(0)
const elevatorPosition = ref(0)
const elevatorSound = ref(null)
const started = ref(false)

const floorRequests = ref([])

function pressButton(dir, floor) {
  if (!floorRequests.value.includes(floor)) {
    floorRequests.value.push(floor)
    floorRequests.value.sort((a, b) => a - b) // keep sorted
  }
}

function moveToFloor(floor) {
  currentFloor.value = floor
  elevatorPosition.value = floor * floorHeight
  if (elevatorSound.value) {
    elevatorSound.value.currentTime = 0
    elevatorSound.value.play().catch(err => console.warn('Autoplay blocked:', err))
  }
}

function simulateElevator() {
  setInterval(() => {
    if (floorRequests.value.length === 0) return;

    const targetFloor = floorRequests.value[0];

    // Move directly to the target floor
    currentFloor.value = targetFloor;
    moveToFloor(currentFloor.value);

    // Remove the floor from the queue once reached
    if (currentFloor.value === targetFloor) {
      floorRequests.value.shift();
    }
  }, 3000); // Keep the interval for animation timing
}

function startElevator() {
  started.value = true
  simulateElevator()
}
</script>

<style scoped>
.container {
  text-align: center;
  padding: 20px;
}

.start-btn {
  padding: 10px 20px;
  font-size: 16px;
  margin: 20px auto;
  cursor: pointer;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 6px;
}

.elevator-shaft {
  position: relative;
  width: 200px;
  height: 500px;
  border: 3px solid #444;
  margin: 30px auto;
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column-reverse;
}

.floor {
  height: 100px;
  border-top: 1px solid #ccc;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 0 10px;
}

.floor-label {
  font-size: 14px;
  color: #333;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.btn {
  background-color: #2196f3;
  color: white;
  border: none;
  padding: 2px 6px;
  border-radius: 4px;
  margin-left: 4px;
  cursor: pointer;
}

.elevator {
  position: absolute;
  left: 20%;
  width: 60%;
  height: 80px;
  background-color: #4caf50;
  border-radius: 8px;
  transition: bottom 1s ease-in-out;
  z-index: 10;
}
</style>

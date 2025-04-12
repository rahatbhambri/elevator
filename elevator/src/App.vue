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
            <button
          v-if="totalFloors - 1 - floor === totalFloors - 1"
          @click="pressButton('down', totalFloors - 1 - floor)"
          :class="{'btn-pressed': buttonStates[totalFloors - 1 - floor].down}"
          class="btn"
        >
          ↓
        </button>

        <button
          v-if="totalFloors - 1 - floor === 0"
          @click="pressButton('up', totalFloors - 1 - floor)"
          :class="{'btn-pressed': buttonStates[totalFloors - 1 - floor].up}"
          class="btn"
        >
          ↑
        </button>

        <button
          v-if="totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1"
          @click="pressButton('up', totalFloors - 1 - floor)"
          :class="{'btn-pressed': buttonStates[totalFloors - 1 - floor].up}"
          class="btn"
        >
          ↑
        </button>
        <button
          v-if="totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1"
          @click="pressButton('down', totalFloors - 1 - floor)"
          :class="{'btn-pressed': buttonStates[totalFloors - 1 - floor].down}"
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
const buttonStates = ref(
  Array.from({ length: totalFloors }, () => ({ up: false, down: false }))
);

function pressButton(dir, floor) {
  if (!floorRequests.value.includes(floor)) {
    floorRequests.value.push(floor);
    floorRequests.value.sort((a, b) => a - b); // Keep sorted
  }

  // Mark the button as pressed
  if (dir === 'up') {
    buttonStates.value[floor].up = true;
  } else if (dir === 'down') {
    buttonStates.value[floor].down = true;
  }
}

function moveToFloor(floor) {
  currentFloor.value = floor;
  elevatorPosition.value = floor * floorHeight;

  // Reset button states for the current floor
  buttonStates.value[floor].up = false;
  buttonStates.value[floor].down = false;

  if (elevatorSound.value) {
    elevatorSound.value.currentTime = 0;
    elevatorSound.value.play().catch(err => console.warn('Autoplay blocked:', err));
  }
}

function hasUpRequestsAbove(currFloor) {
  return floorRequests.value.some(floor => floor > currFloor);
}

function hasDownRequestsBelow(currFloor) {
  return floorRequests.value.some(floor => floor < currFloor);
}

function getCurrentDirection() {
  // Determine the current direction based on the next request in the queue
  if (floorRequests.value.length === 0) return 0; // No direction
  const nextFloor = floorRequests.value[0];
  return nextFloor > currentFloor.value ? 1 : -1; // 1 for up, -1 for down
}

function simulateElevator() {
  setInterval(() => {
    if (floorRequests.value.length === 0) return;

    const dir = getCurrentDirection(); // Get the current direction of the elevator
    const currFloor = currentFloor.value;

    if (dir === 1) { // Moving up
      if (hasUpRequestsAbove(currFloor)) {
        // Serve all "up" requests above the current floor
        for (let i = currFloor + 1; i < totalFloors; i++) {
          if (floorRequests.value.includes(i)) {
            moveToFloor(i);
            floorRequests.value = floorRequests.value.filter(f => f !== i); // Remove served floor
          }
        }
      } else if (hasDownRequestsBelow(currFloor)) {
        // Switch to "down" direction and serve all "down" requests
        for (let i = currFloor - 1; i >= 0; i--) {
          if (floorRequests.value.includes(i)) {
            moveToFloor(i);
            floorRequests.value = floorRequests.value.filter(f => f !== i); // Remove served floor
          }
        }
      }
    } else if (dir === -1) { // Moving down
      if (hasDownRequestsBelow(currFloor)) {
        // Serve all "down" requests below the current floor
        for (let i = currFloor - 1; i >= 0; i--) {
          if (floorRequests.value.includes(i)) {
            moveToFloor(i);
            floorRequests.value = floorRequests.value.filter(f => f !== i); // Remove served floor
          }
        }
      } else if (hasUpRequestsAbove(currFloor)) {
        // Switch to "up" direction and serve all "up" requests
        for (let i = currFloor + 1; i < totalFloors; i++) {
          if (floorRequests.value.includes(i)) {
            moveToFloor(i);
            floorRequests.value = floorRequests.value.filter(f => f !== i); // Remove served floor
          }
        }
      }
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
.btn.btn-pressed {
  background-color: grey !important; /* Ensure it overrides other styles */
  cursor: not-allowed;
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

package elevator

getcurrfloor()int{
}

getcurrdir() int {

}


next_floor_queue = []
ups = [True False False ..]
downs = [True False ..]



GetNextMove(dir int, curr_fl int){
	// momentum first 
	if dir ==1{
		if or(ups[curr_fl:]){
			// Go up till last up is True stopping at all set floors
			for i:= curr_fl+1; i<n ; i +=1{
				if ups[i]{
					next_floor_queue += [i]
				}
			}
		}else{
			// Go down till last down is True	
		}
		
	}else if dir == -1{
		// look for downest active floor

		// else look for uppest active floor
	}
}


pressButton(button string, floor int){
	if button == "up"{
		ups[floor] = True
	}else{

	}
}

// Elevator
func move(from_floor int, to_floor int) (bool, error){
	if !lockAll(){
		return false, errors.New("Door lock error")
	}
	// movement
	unlock(to_floor)
}

func lockAll()


package task

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BusInfo struct {
	NumFamilies   int
	NumPassengers int
}

type BusInfoTemporary struct {
	IsTaken       bool
	NumPassengers int
}

func CalculateMinBuses(reader *bufio.Reader) {
	fmt.Print("Enter the number of families: ")
	inputKeluarga, _ := reader.ReadString('\n')
	inputKeluarga = strings.TrimSpace(inputKeluarga)
	n, err := strconv.Atoi(inputKeluarga)
	if err != nil || n <= 0 {
		fmt.Println("Invalid family number input")
		return
	}

	fmt.Print("Enter the number of members in the family (separated by spaces): ")
	inputAnggota, _ := reader.ReadString('\n')
	inputAnggota = strings.TrimSpace(inputAnggota)

	anggotaStrings := strings.Split(inputAnggota, " ")
	var familySizes []int

	for _, str := range anggotaStrings {
		size, err := strconv.Atoi(str)
		if err != nil || size <= 0 {
			fmt.Println("Invalid family member input")
			return
		}
		familySizes = append(familySizes, size)
	}

	if len(familySizes) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	minBuses := calculateMinBusesHelper(familySizes)
	fmt.Printf("Minimum bus required is : %d\n", minBuses)
}

func calculateMinBusesHelper(familyMembersList []int) (totalBuses int) {
	const busCapacity = 4
	const maxFamiliesPerBus = 2

	var busesInfoTmp []*BusInfoTemporary
	for _, numFamilyMembers := range familyMembersList {
		busesInfoTmp = append(busesInfoTmp, &BusInfoTemporary{NumPassengers: numFamilyMembers, IsTaken: false})
	}

	var busesInfo []*BusInfo

	for _, busInfoTmp := range busesInfoTmp {
		if busInfoTmp.IsTaken {
			continue
		}

		busInfoTmp.IsTaken = true

		if len(busesInfo) == 0 {
			if busInfoTmp.NumPassengers < busCapacity {

				busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: busInfoTmp.NumPassengers})
				continue
			}

			numBuses := busInfoTmp.NumPassengers / busCapacity
			remainingPassengers := busInfoTmp.NumPassengers % busCapacity

			for i := 0; i <= numBuses; i++ {
				if i == numBuses && remainingPassengers != 0 {
					busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: remainingPassengers})
					break
				}

				if i == numBuses && remainingPassengers == 0 {
					break
				}

				busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: busCapacity})
			}

			continue
		}

		if busInfoTmp.NumPassengers == busCapacity {
			busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: busCapacity})
			continue
		}

		stop := false

		if busInfoTmp.NumPassengers < busCapacity {
			for _, busInfoTmp2 := range busesInfoTmp {
				if !busInfoTmp2.IsTaken && (busInfoTmp2.NumPassengers+busInfoTmp.NumPassengers) == busCapacity {

					busesInfo = append(busesInfo, &BusInfo{NumFamilies: maxFamiliesPerBus, NumPassengers: busCapacity})

					busInfoTmp2.IsTaken = true
					stop = true
					break
				}
			}

			if stop {
				continue
			}

			for _, busInfo := range busesInfo {
				if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+busInfoTmp.NumPassengers) == busCapacity {

					busInfo.NumFamilies = maxFamiliesPerBus
					busInfo.NumPassengers = busCapacity

					stop = true
					break
				}
			}

			if stop {
				continue
			}

			for _, busInfo := range busesInfo {

				if busInfo.NumFamilies < maxFamiliesPerBus && busInfo.NumPassengers != busCapacity && (busInfo.NumPassengers+busInfoTmp.NumPassengers) > busCapacity {

					remainingPassengers := (busInfo.NumPassengers + busInfoTmp.NumPassengers) - busCapacity

					busInfo.NumFamilies = maxFamiliesPerBus
					busInfo.NumPassengers = busCapacity

					for _, busInfoTmp := range busesInfoTmp {
						if !busInfoTmp.IsTaken && (busInfoTmp.NumPassengers+remainingPassengers) == busCapacity {
							busesInfo = append(busesInfo, &BusInfo{NumFamilies: maxFamiliesPerBus, NumPassengers: busCapacity})

							busInfoTmp.IsTaken = true
							stop = true
							remainingPassengers = 0
							break
						}
					}

					if remainingPassengers != 0 {
						for _, busInfo := range busesInfo {
							if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+remainingPassengers) == busCapacity {
								busInfo.NumFamilies = maxFamiliesPerBus
								busInfo.NumPassengers = busCapacity

								remainingPassengers = 0
								stop = true
								break
							}
						}
					}

					if remainingPassengers != 0 {
						busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: remainingPassengers})
						remainingPassengers = 0
						stop = true
						break
					}
				}

				if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+busInfoTmp.NumPassengers) < busCapacity {
					busInfo.NumFamilies = maxFamiliesPerBus
					busInfo.NumPassengers = busInfo.NumPassengers + busInfoTmp.NumPassengers

					stop = true
					break
				}
			}

			if stop {
				continue
			}

			busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: busInfoTmp.NumPassengers})
			stop = true
		}

		if busInfoTmp.NumPassengers > busCapacity {
			numBuses := busInfoTmp.NumPassengers / busCapacity
			remainingPassengers := busInfoTmp.NumPassengers % busCapacity

			for i := 0; i <= numBuses; i++ {
				if i == numBuses {
					if remainingPassengers != 0 {

						for _, busInfoTmp := range busesInfoTmp {
							if !busInfoTmp.IsTaken && (busInfoTmp.NumPassengers+remainingPassengers) == busCapacity {
								busesInfo = append(busesInfo, &BusInfo{NumFamilies: maxFamiliesPerBus, NumPassengers: busCapacity})

								remainingPassengers = 0
								stop = true
								break
							}
						}

						if remainingPassengers != 0 {
							for _, busInfo := range busesInfo {
								if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+remainingPassengers) == busCapacity {
									busInfo.NumFamilies = maxFamiliesPerBus
									busInfo.NumPassengers = busCapacity

									remainingPassengers = 0
									stop = true
									break
								}
							}
						}

						if remainingPassengers != 0 {
							for _, busInfo := range busesInfo {
								if busInfo.NumFamilies < maxFamiliesPerBus && busInfo.NumPassengers != busCapacity && (busInfo.NumPassengers+remainingPassengers) > busCapacity {
									remainingPassengers := (busInfo.NumPassengers + busInfoTmp.NumPassengers) - busCapacity

									busInfo.NumFamilies = maxFamiliesPerBus
									busInfo.NumPassengers = busCapacity

									for _, busInfoTmp := range busesInfoTmp {
										if !busInfoTmp.IsTaken && (busInfoTmp.NumPassengers+remainingPassengers) == busCapacity {
											busesInfo = append(busesInfo, &BusInfo{NumFamilies: maxFamiliesPerBus, NumPassengers: busCapacity})

											busInfoTmp.IsTaken = true
											stop = true
											remainingPassengers = 0
											break
										}
									}

									if remainingPassengers != 0 {
										for _, busInfo := range busesInfo {
											if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+remainingPassengers) == busCapacity {
												busInfo.NumFamilies = maxFamiliesPerBus
												busInfo.NumPassengers = busCapacity

												remainingPassengers = 0
												stop = true
												break
											}
										}
									}

									if remainingPassengers != 0 {
										busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: remainingPassengers})
										remainingPassengers = 0
										stop = true
										break
									}
								}

								if busInfo.NumFamilies < maxFamiliesPerBus && (busInfo.NumPassengers+remainingPassengers) < busCapacity {
									busInfo.NumFamilies = maxFamiliesPerBus
									busInfo.NumPassengers += remainingPassengers

									stop = true
									remainingPassengers = 0
									break
								}
							}
						}

						if remainingPassengers != 0 {
							busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: remainingPassengers})
						}
					}

					if remainingPassengers == 0 {
						break
					}
				}

				busesInfo = append(busesInfo, &BusInfo{NumFamilies: 1, NumPassengers: busCapacity})
			}
		}
	}

	return len(busesInfo)
}

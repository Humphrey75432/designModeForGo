package medium

type StationManager struct {
	IsPlatformFree bool
	TrainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
	}
}

func (s *StationManager) CanArrive(train Train) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, train)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainQueue.PermitArrival()
	}
}

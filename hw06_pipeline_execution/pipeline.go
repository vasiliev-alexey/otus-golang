package hw06_pipeline_execution //nolint:golint,stylecheck
import "log"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outCh := in
	for _, stage := range stages {
		outCh = func(in In) (out Out) {
			bindCh := make(Bi)

			go func() {
				defer close(bindCh)

				for {
					select {
					case <-done:
						log.Println("graceful shutdown processing")
						return
					case v, ok := <-in:
						if !ok {
							return
						}
						select {
						case <-done:
						case bindCh <- v:
							log.Println("перекладываем значение в связующий канал")
						}
					}
				}
			}()
			log.Println("Вызываем трансформер на связующем канале")
			return stage(bindCh)
		}(outCh)
	}

	return outCh
}

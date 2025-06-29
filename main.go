package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	windowWidth  = 1000
	windowHeight = 600
	charWidth    = 12
	charHeight   = 18
	cols         = windowWidth / charWidth   // 83 columns
	rows         = windowHeight / charHeight // 33 rows
	charSet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()_+-=[]{}|;:,.<>?/~`"
)

type MatrixChar struct {
	text     *canvas.Text
	char     rune
	age      int
	maxAge   int
	isActive bool
}



type MatrixColumn struct {
	chars       []*MatrixChar
	dropY       int
	speed       int
	height      int
	llmName     string
	llmActive   bool
	llmStartPos int
}

type MatrixScreen struct {
	columns      []*MatrixColumn
	content      *fyne.Container
	nextLLMTime  time.Time
}

func NewMatrixScreen() *MatrixScreen {
	m := &MatrixScreen{
		columns:     make([]*MatrixColumn, cols),
		content:     container.NewWithoutLayout(),
		nextLLMTime: time.Now().Add(1 * time.Second), // First LLM name appears after 1 second
	}

	// Initialize columns
	for i := 0; i < cols; i++ {
		m.columns[i] = &MatrixColumn{
			chars:       make([]*MatrixChar, rows),
			dropY:       -rand.Intn(rows),
			speed:       rand.Intn(3) + 1,
			height:      rand.Intn(15) + 10,
			llmName:     "",
			llmActive:   false,
			llmStartPos: 0,
		}

		// Initialize characters in each column
		for j := 0; j < rows; j++ {
			char := &MatrixChar{
				text:     canvas.NewText(string(getRandomChar()), color.RGBA{0, 255, 0, 255}),
				char:     getRandomChar(),
				age:      0,
				maxAge:   rand.Intn(20) + 10,
				isActive: false,
			}
			char.text.TextSize = 14
			char.text.Alignment = fyne.TextAlignCenter
			char.text.Move(fyne.NewPos(float32(i*charWidth), float32(j*charHeight)))
			m.columns[i].chars[j] = char
			m.content.Add(char.text)
		}
	}

	return m
}

func getRandomChar() rune {
	return rune(charSet[rand.Intn(len(charSet))])
}

func getLLMNames() []string {
	return []string{
		"GPT-4", "CLAUDE", "GEMINI", "LLAMA", "MISTRAL",
		"ANTHROPIC", "OPENAI", "COHERE", "PALM", "BARD",
		"CHATGPT", "CODEX", "DAVINCI", "ADA", "CURIE",
		"ALPACA", "VICUNA", "FALCON", "MPT", "DOLLY",
		"STABLELM", "PYTHIA", "BLOOM", "OPT", "T5",
	}
}

func (m *MatrixScreen) startLLMColumn() {
	names := getLLMNames()
	selectedName := names[rand.Intn(len(names))]

	// Find a column that's not currently showing an LLM name
	availableColumns := make([]int, 0)
	for i, col := range m.columns {
		if !col.llmActive {
			availableColumns = append(availableColumns, i)
		}
	}

	if len(availableColumns) > 0 {
		colIndex := availableColumns[rand.Intn(len(availableColumns))]
		col := m.columns[colIndex]

		col.llmName = selectedName
		col.llmActive = true
		col.llmStartPos = col.dropY
		col.dropY = -5 // Start just above screen
		col.height = len(selectedName) + 5 // Ensure stream is long enough for the name

		m.nextLLMTime = time.Now().Add(time.Duration(rand.Intn(1500)+500) * time.Millisecond) // 0.5-2 seconds

		// println("LLM name starting to fall:", selectedName, "at column", colIndex) // Debug message
	}
}

func (m *MatrixScreen) Update() {
	now := time.Now()

	// Check if it's time to start a new LLM column
	if now.After(m.nextLLMTime) {
		m.startLLMColumn()
	}

	for _, col := range m.columns {
		// Move the drop down
		col.dropY++

		// Check if LLM name has finished falling
		if col.llmActive && col.dropY > rows+len(col.llmName) {
			col.llmActive = false
			col.llmName = ""
		}

		// Reset column if it's gone too far
		if col.dropY > rows+col.height {
			col.dropY = -rand.Intn(rows)
			col.speed = rand.Intn(3) + 1
			col.height = rand.Intn(15) + 10
			col.llmActive = false // Reset LLM name when column resets
			col.llmName = ""
		}

										// Update characters in the column
		for rowIndex, char := range col.chars {
			// Check if this character should be part of the falling stream
			inStream := rowIndex >= col.dropY && rowIndex < col.dropY+col.height

			if inStream {
				if !char.isActive {
					// Start new character
					char.isActive = true
					char.age = 0
					char.char = getRandomChar()
					char.text.Text = string(char.char)
				}

				// Check if this should be an LLM name character
				if col.llmActive {
					nameIndex := rowIndex - col.dropY
					if nameIndex >= 0 && nameIndex < len(col.llmName) {
						// This position should show an LLM name character
						char.char = rune(col.llmName[nameIndex])
						char.text.Text = string(char.char)
						char.text.Color = color.RGBA{255, 255, 100, 255} // Bright yellow-white for LLM names
					} else {
						// Regular Matrix character in LLM column
						distanceFromHead := rowIndex - col.dropY
						if distanceFromHead == 0 {
							char.text.Color = color.RGBA{255, 255, 255, 255} // Head white
						} else {
							alpha := 255 - (distanceFromHead * 15)
							if alpha < 50 {
								alpha = 50
							}
							char.text.Color = color.RGBA{0, uint8(alpha), 0, 255} // Green fade
						}
						if rand.Float32() < 0.1 {
							char.char = getRandomChar()
							char.text.Text = string(char.char)
						}
					}
				} else {
					// Update character appearance based on position in stream
					distanceFromHead := rowIndex - col.dropY
					alpha := 255

					if distanceFromHead == 0 {
						// Head of the stream - bright white
						char.text.Color = color.RGBA{255, 255, 255, 255}
					} else {
						// Body of the stream - fade from bright green to dark green
						alpha = 255 - (distanceFromHead * 15)
						if alpha < 50 {
							alpha = 50
						}
						char.text.Color = color.RGBA{0, uint8(alpha), 0, 255}
					}

					// Occasionally change character
					if rand.Float32() < 0.1 {
						char.char = getRandomChar()
						char.text.Text = string(char.char)
					}
				}

			} else if char.isActive {
				// Character is fading out
				char.age++
				fadeAlpha := 255 - (char.age * 20)
				if fadeAlpha <= 0 {
					char.isActive = false
					char.text.Text = " "
					char.text.Color = color.RGBA{0, 0, 0, 0}
				} else {
					char.text.Color = color.RGBA{0, uint8(fadeAlpha/4), 0, uint8(fadeAlpha)}
				}
			}

			char.text.Refresh()
		}
	}
}

func (m *MatrixScreen) GetContent() *fyne.Container {
	return m.content
}

func main() {
	rand.Seed(time.Now().UnixNano())

	myApp := app.New()
	myWindow := myApp.NewWindow("Matrix")
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(windowWidth, windowHeight))

	// Create black background
	bg := canvas.NewRectangle(color.RGBA{0, 0, 0, 255})
	bg.Resize(fyne.NewSize(windowWidth, windowHeight))

	matrix := NewMatrixScreen()

	// Create main container with background
	content := container.NewWithoutLayout(bg, matrix.GetContent())
	myWindow.SetContent(content)

	// Start animation loop
	go func() {
		ticker := time.NewTicker(120 * time.Millisecond) // Slightly slower for better visibility
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				matrix.Update()
			}
		}
	}()

	// Add instructions
	instructions := widget.NewLabel("Press Ctrl+C to exit")
	instructions.Move(fyne.NewPos(10, 10))
	instructions.Resize(fyne.NewSize(200, 20))
	content.Add(instructions)

	myWindow.ShowAndRun()
}
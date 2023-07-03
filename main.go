

package main

import (
	"context"
	"fmt"
	

	openai "github.com/sashabaranov/go-openai"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	c := openai.NewClient("YOUR_OPENAI_API_KEY")
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "./YOUR_MP3_FILE_HERE.mp3",
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}

	err = saveTranscriptionToPDF(resp.Text, "THE_PDF_TRANSCRIPT_OUTPUT_HERE.pdf")
	if err != nil {
		fmt.Printf("PDF creation error: %v\n", err)
		return
	}
	fmt.Println("Transcription saved")
}





func saveTranscriptionToPDF(transcription, filePath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("arial", "B", 16)
	pdf.MultiCell(0, 12, transcription, "", "", false)
	return pdf.OutputFileAndClose(filePath)
}



**This is an open source Go project that converts an MP3 file into a PDF file, leveraging the OpenAI Whisper API for audio transcription.** <br> 

The resulting transcription is saved as a PDF document for easy reference and sharing.


**Installation** <br> 
To use this project, you will need to have Go installed on your machine. You can download and install Go from the official website: https://golang.org

Next, clone the repository from GitHub:
`git clone https://github.com/laurentknauss/mp3-to-pdf-converter.git`


Navigate to the project directory:
`cd mp3-to-pdf-converter`

Install the project dependencies:
`go mod download`   <br>


**Usage**

Before running the code, make sure you have an OpenAI API key. If you don't have one, you can sign up and obtain an API key from the OpenAI website: https://openai.com

Replace "YOUR_OPENAI_API_KEY" in the code with your actual OpenAI API key.


`c := openai.NewClient("YOUR_OPENAI_API_KEY")` <br> 
<br>


Replace "YOUR_MP3_FILE_HERE.mp3" with the path to your input MP3 file.

`req := openai.AudioRequest{
    Model:    openai.Whisper1,
    FilePath: "./YOUR_MP3_FILE_HERE.mp3",
}`  <br>
<br>



Replace "THE_PDF_TRANSCRIPT_OUTPUT_HERE.pdf" with the desired path and name for the output PDF file.


`err = saveTranscriptionToPDF(resp.Text, "THE_PDF_TRANSCRIPT_OUTPUT_HERE.pdf")` <br>
<br> 


Once you have made these changes, you can run the code:

`go run main.go`



The transcription will be generated and saved as a PDF file in the specified location.

**Contributing** <br>  
Contributions to this project are welcome. If you find any issues or have suggestions for improvements, please open an issue on the GitHub repository: https://github.com/your-username/your-repo
You can also submit pull requests with your proposed changes. We appreciate your help in making this project better.


**License** 

This project is licensed under the MIT License. Feel free to use and modify the code according to your needs.

**Acknowledgements** <br>

This project uses the following third-party libraries:

*go-openai - Go client library for the OpenAI API* <br>
*gofpdf - PDF document generation library for Go*

I thank the authors and contributors of these libraries for their valuable work.

**Disclaimer** <br> 

This project is provided as-is with no warranty or guarantee of its functionality. Use it at your own risk. We are not responsible for any consequences resulting from the use of this project.












































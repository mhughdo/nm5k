package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"github.com/zellyn/kooky"
)

const DefaultEditor = "nano"

type PreferredEditorResolver func() string

// GetPreferredEditorFromEnvironment returns the user's editor as defined by the
// `$EDITOR` environment variable, or the `DefaultEditor` if it is not set.
func GetPreferredEditorFromEnvironment() string {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		return DefaultEditor
	}

	return editor
}

func resolveEditorArguments(executable string, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") {
		args = append([]string{"--wait"}, args...)
	}

	// Other common editors

	return args
}

// OpenFileInEditor opens filename in a text editor.
func OpenFileInEditor(filename string, params ...PreferredEditorResolver) error {
	// Get the full executable path for the editor.
	var resolveEditor PreferredEditorResolver
	if len(params) > 0 {
		resolveEditor = params[0]
	} else {
		resolveEditor = GetPreferredEditorFromEnvironment
	}

	executable, err := exec.LookPath(resolveEditor())
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, resolveEditorArguments(executable, filename)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CaptureInputFromEditor opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func CaptureInputFromEditor(defaultMessage string) ([]byte, error) {
	message := []byte(defaultMessage)

	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	_, err = file.Write(message)

	if err != nil {
		log.Fatalln(err)
	}

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

// GetCookie get cookie automatically
func GetCookie() (string, error) {
	domain := "www.chatwork.com"
	cookieName := "cwssid"
	var cookies []*kooky.Cookie
	var err error

	usr, _ := user.Current()

	cookiesFile := fmt.Sprintf("%s/Library/Cookies/Cookies.binarycookies", usr.HomeDir)
	cookies, err = kooky.ReadSafariCookies(cookiesFile, domain, "", time.Time{})
	if len(cookies) == 0 {
		// safari had none, try chrome
		cookiesFile := fmt.Sprintf("%s/Library/Application Support/Google/Chrome/Default/Cookies", usr.HomeDir)
		cookies, err = kooky.ReadChromeCookies(cookiesFile, domain, "", time.Time{})
	}

	if err != nil {
		return "", err
	}

	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			return cookie.Value, nil
		}
	}

	return "", nil
}

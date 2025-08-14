package gui

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/ygzaydn/go-dnswatcher/internal/config"
	"github.com/ygzaydn/go-dnswatcher/internal/eventlog"
	"github.com/ygzaydn/go-dnswatcher/internal/kpi"
)

func Start() {
	a := app.New()
	w := a.NewWindow("DNSWatcher")

	// Summary Tab
	summaryLabel := widget.NewLabel("Welcome to DNSWatcher!")
	summaryStatsBind := binding.NewString()
	summaryContent := widget.NewLabelWithData(summaryStatsBind)
	summaryTab := container.NewVBox(summaryLabel, summaryContent)

	// Event log tab
	eventLogLabel := widget.NewLabel("Event Log")
	eventLogStatsBind := binding.NewString()
	eventLogContent := widget.NewLabelWithData(eventLogStatsBind)
	scrollableEventLog := container.NewVScroll(eventLogContent)
	scrollableEventLog.SetMinSize(fyne.NewSize(1260, 600))
	eventLogTab := container.NewVBox(eventLogLabel, scrollableEventLog)

	// Config tab
	configLabel := widget.NewLabel("Configuration")
	cfg, err := config.LoadConfig("../config.yaml")
	var configContent *widget.Label
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		configContent = widget.NewLabel("Failed to load configuration.")
	} else {
		configContent = widget.NewLabel(config.ParseConfig(cfg))
	}
	configTab := container.NewVBox(configLabel, configContent)

	tabs := container.NewAppTabs(
		container.NewTabItem("Summary", summaryTab),
		container.NewTabItem("Event Log", eventLogTab),
		container.NewTabItem("Configuration", configTab),
	)

	w.Resize(fyne.NewSize(1260, 800))

	w.SetContent(
		tabs,
	)

	go func() {
		for {
			time.Sleep(time.Second)
			metrics := kpi.GetMetrics()
			stats := kpi.StatsString(metrics)
			_ = summaryStatsBind.Set(stats)
			events := eventlog.GetAll()
			_ = eventLogStatsBind.Set(strings.Join(events, "\n"))
		}
	}()

	w.ShowAndRun()
}

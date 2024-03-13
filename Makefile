include config.mk

install:
	@echo "Installing csv-viewer script..."
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 csv-viewer $(DESTDIR)$(BINDIR)
	@echo "Installing man page..."
	install -d $(DESTDIR)$(MANDIR)
	install -m 644 csv-viewer.1 $(DESTDIR)$(MANDIR)
	@echo "Installation complete."

uninstall:
	@echo "Uninstalling csv-viewer script..."
	rm -f $(DESTDIR)$(BINDIR)/csv-viewer
	@echo "Removing man page..."
	rm -f $(DESTDIR)$(MANDIR)/csv-viewer.1
	@echo "Uninstallation complete."

.PHONY: install uninstall
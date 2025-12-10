#!/bin/bash
set -e

# Input/Output
TEMPLATE="readme.in.md"
OUTPUT="README.md"

# AUTO-Block Inhalte definieren
# Du kannst hier mehrere Blöcke definieren, z.B. API, BUILD, etc.
declare -A AUTO_BLOCKS

# Beispiel: ein Block "AUTO"
# Datei: auto.txt
AUTO_BLOCKS["AUTO"]="$(<auto.md)"

# --- Script ---
cp "$TEMPLATE" "$OUTPUT"

for BLOCK in "${!AUTO_BLOCKS[@]}"; do
    CONTENT="${AUTO_BLOCKS[$BLOCK]}"

    # awk ersetzt den jeweiligen AUTO-Block
    awk -v content="$CONTENT" -v block="$BLOCK" '
    BEGIN {inblock=0}
    {
        if ($0 ~ "<!-- BEGIN " block " -->") {
            print "<!-- BEGIN " block " -->"
            print content
            inblock=1
            next
        }
        if ($0 ~ "<!-- END " block " -->") {
            inblock=0
            print "<!-- END " block " -->"
            next
        }
        if (!inblock) print
    }' "$OUTPUT" > "${OUTPUT}.tmp"

    mv "${OUTPUT}.tmp" "$OUTPUT"
done

echo "README.md erfolgreich generiert."
